package detector

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	discoveryv1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/utils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/utils/workloadutils"
	"github.com/solo-io/go-utils/contextutils"
	sets2 "github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/pointer"
)

const (
	// TODO: allow for specifying specific meshes.
	// Currently this annotation assumes that there is only one mesh per cluster, and therefore the corresponding
	// TrafficTarget will be associated with that mesh.
	DiscoveryMeshAnnotation = "discovery.mesh.gloo.solo.io/enabled"
)

var (
	skippedLabels = sets.NewString(
		"pod-template-hash",
		"service.istio.io/canonical-revision",
	)
)

// the TrafficTargetDetector detects TrafficTargets from services
// whose backing pods are injected with a Mesh sidecar.
// If no Mesh is detected, nil is returned
type TrafficTargetDetector interface {
	DetectTrafficTarget(
		ctx context.Context,
		service *corev1.Service,
		workloads discoveryv1alpha2sets.WorkloadSet,
		meshes discoveryv1alpha2sets.MeshSet,
	) *v1alpha2.TrafficTarget
}

type trafficTargetDetector struct{}

func NewTrafficTargetDetector() TrafficTargetDetector {
	return &trafficTargetDetector{}
}

func (t *trafficTargetDetector) DetectTrafficTarget(
	ctx context.Context,
	service *corev1.Service,
	meshWorkloads discoveryv1alpha2sets.WorkloadSet,
	meshes discoveryv1alpha2sets.MeshSet,
) *v1alpha2.TrafficTarget {

	kubeService := &v1alpha2.TrafficTargetSpec_KubeService{
		Ref:                    ezkube.MakeClusterObjectRef(service),
		WorkloadSelectorLabels: service.Spec.Selector,
		Labels:                 service.Labels,
		Ports:                  convertPorts(service),
	}

	trafficTarget := &v1alpha2.TrafficTarget{
		ObjectMeta: utils.DiscoveredObjectMeta(service),
		Spec: v1alpha2.TrafficTargetSpec{
			Type: &v1alpha2.TrafficTargetSpec_KubeService_{
				KubeService: kubeService,
			},
		},
	}

	// If the service is not associated with a mesh, do not create a traffic target
	validTrafficTarget := addMeshForKubeServiceIfValid(
		ctx,
		trafficTarget,
		service,
		kubeService,
		meshWorkloads,
		meshes,
	)
	if !validTrafficTarget {
		return nil
	}

	return trafficTarget
}

func addMeshForKubeServiceIfValid(
	ctx context.Context,
	tt *v1alpha2.TrafficTarget,
	service *corev1.Service,
	kubeService *v1alpha2.TrafficTargetSpec_KubeService,
	meshWorkloads discoveryv1alpha2sets.WorkloadSet,
	meshes discoveryv1alpha2sets.MeshSet,
) bool {

	var validMesh *v1.ObjectRef

	// TODO: support subsets from services which have been discovered via the annotation
	discoveryEnabled, ok := service.Annotations[DiscoveryMeshAnnotation]
	if ok && discoveryEnabled == "true" {

		// Search for mesh which exists on the same cluster as the annotated service
		for _, mesh := range meshes.List() {
			mesh := mesh
			switch typedMesh := mesh.Spec.GetMeshType().(type) {
			case *v1alpha2.MeshSpec_Osm:
				if typedMesh.Osm.GetInstallation().GetCluster() == service.GetClusterName() {
					validMesh = ezkube.MakeObjectRef(mesh)
					break
				}
			case *v1alpha2.MeshSpec_Istio_:
				if typedMesh.Istio.GetInstallation().GetCluster() == service.GetClusterName() {
					validMesh = ezkube.MakeObjectRef(mesh)
					break
				}
			}
		}

		if validMesh == nil {
			contextutils.LoggerFrom(ctx).Errorf(
				"mesh could not be found for annotated service %s",
				sets2.TypedKey(service),
			)
		}
	}

	if validMesh != nil {
		tt.Spec.Mesh = validMesh
		return true
	}

	// if no mesh was found from the annotation, check the workloads
	backingWorkloads := workloadutils.FindBackingWorkloads(kubeService, meshWorkloads)
	// if discovery is enabled, do not return
	if backingWorkloads.Length() == 0 {
		return false
	}
	handleWorkloadDiscoveredMesh(ctx, tt, backingWorkloads, kubeService)
	return true
}

func handleWorkloadDiscoveredMesh(
	ctx context.Context,
	tt *v1alpha2.TrafficTarget,
	backingWorkloads discoveryv1alpha2sets.WorkloadSet,
	kubeService *v1alpha2.TrafficTargetSpec_KubeService,
) {
	// all backing workloads should be in the same mesh
	validMesh := backingWorkloads.List()[0].Spec.Mesh

	// derive subsets from backing workloads
	kubeService.Subsets = findSubsets(backingWorkloads)

	// Add a refrence to each workload to the TrafficTarget
	for _, workload := range backingWorkloads.List() {
		tt.Spec.Workloads = append(tt.Spec.Workloads, &v1.ObjectRef{
			Name:      workload.GetName(),
			Namespace: workload.GetNamespace(),
		})
	}

	tt.Spec.Mesh = validMesh
}

// expects a list of just the workloads that back the service you're finding subsets for
func findSubsets(backingWorkloads discoveryv1alpha2sets.WorkloadSet) map[string]*v1alpha2.TrafficTargetSpec_KubeService_Subset {
	uniqueLabels := make(map[string]sets.String)
	for _, backingWorkload := range backingWorkloads.List() {
		for key, val := range backingWorkload.Spec.GetKubernetes().GetPodLabels() {
			// skip known kubernetes values
			if skippedLabels.Has(key) {
				continue
			}
			existing, ok := uniqueLabels[key]
			if !ok {
				uniqueLabels[key] = sets.NewString(val)
			} else {
				existing.Insert(val)
			}
		}
	}
	/*
		Only select the keys with > 1 value
		The subsets worth noting will be sets of labels which share the same key, but have different values, such as:

			version:
				- v1
				- v2
	*/
	subsets := make(map[string]*v1alpha2.TrafficTargetSpec_KubeService_Subset)
	for k, v := range uniqueLabels {
		if v.Len() > 1 {
			subsets[k] = &v1alpha2.TrafficTargetSpec_KubeService_Subset{Values: v.List()}
		}
	}
	if len(subsets) == 0 {
		// important to return nil instead of empty map for asserting equality
		return nil
	}
	return subsets
}

func convertPorts(service *corev1.Service) (ports []*v1alpha2.KubeServicePort) {
	for _, kubePort := range service.Spec.Ports {
		ports = append(ports, &v1alpha2.KubeServicePort{
			Port:        uint32(kubePort.Port),
			Name:        kubePort.Name,
			Protocol:    string(kubePort.Protocol),
			AppProtocol: pointer.StringPtrDerefOr(kubePort.AppProtocol, ""),
		})
	}
	return ports
}
