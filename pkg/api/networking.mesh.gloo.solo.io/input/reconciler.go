// Code generated by skv2. DO NOT EDIT.

package input

import (
	"context"
	"time"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	appmesh_k8s_aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	appmesh_k8s_aws_v1beta2_controllers "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/controller"
	networking_istio_io_v1alpha3_controllers "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/controller"
	security_istio_io_v1beta1_controllers "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/controller"
	v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	certificates_mesh_gloo_solo_io_v1alpha2_controllers "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/controller"
	discovery_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	discovery_mesh_gloo_solo_io_v1alpha2_controllers "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/controller"
	networking_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	networking_mesh_gloo_solo_io_v1alpha2_controllers "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/controller"
	settings_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	settings_mesh_gloo_solo_io_v1alpha2_controllers "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/controller"
	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1"
	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_controllers "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1alpha1/controller"
	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_controllers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/controller"
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// The Input Reconciler calls a simple func(id) error whenever a
// storage event is received for any of:
// * CertificatesMeshGlooSoloIov1Alpha2IssuedCertificates
// * CertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives
// * XdsAgentEnterpriseMeshGlooSoloIov1Alpha1XdsConfigs
// * AppmeshK8SAwsv1Beta2VirtualServices
// * AppmeshK8SAwsv1Beta2VirtualNodes
// * AppmeshK8SAwsv1Beta2VirtualRouters
// * NetworkingIstioIov1Alpha3DestinationRules
// * NetworkingIstioIov1Alpha3EnvoyFilters
// * NetworkingIstioIov1Alpha3Gateways
// * NetworkingIstioIov1Alpha3ServiceEntries
// * NetworkingIstioIov1Alpha3VirtualServices
// * SecurityIstioIov1Beta1AuthorizationPolicies
// * V1ConfigMaps
// from a remote cluster.
// * SettingsMeshGlooSoloIov1Alpha2Settings
// * DiscoveryMeshGlooSoloIov1Alpha2TrafficTargets
// * DiscoveryMeshGlooSoloIov1Alpha2Workloads
// * DiscoveryMeshGlooSoloIov1Alpha2Meshes
// * NetworkingMeshGlooSoloIov1Alpha2TrafficPolicies
// * NetworkingMeshGlooSoloIov1Alpha2AccessPolicies
// * NetworkingMeshGlooSoloIov1Alpha2VirtualMeshes
// * NetworkingMeshGlooSoloIov1Alpha2FailoverServices
// * V1Secrets
// * MulticlusterSoloIov1Alpha1KubernetesClusters
// from the local cluster.

type ReconcileOptions struct {
	Remote RemoteReconcileOptions
	Local  LocalReconcileOptions

	// the ReconcileInterval, if greater than 0, will limit the number of reconciles
	// to one per interval.
	ReconcileInterval time.Duration
}

// register the given multi cluster reconcile func with the cluster watcher
// register the given single cluster reconcile func with the local manager
func RegisterInputReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	multiClusterReconcileFunc input.MultiClusterReconcileFunc,
	mgr manager.Manager,
	singleClusterReconcileFunc input.SingleClusterReconcileFunc,
	options ReconcileOptions,
) (input.InputReconciler, error) {

	base := input.NewInputReconciler(
		ctx,
		multiClusterReconcileFunc,
		singleClusterReconcileFunc,
		options.ReconcileInterval,
	)

	// initialize reconcile loops

	// initialize CertificatesMeshGlooSoloIov1Alpha2IssuedCertificates reconcile loop for remote clusters
	certificates_mesh_gloo_solo_io_v1alpha2_controllers.NewMulticlusterIssuedCertificateReconcileLoop("IssuedCertificate", clusters, options.Remote.CertificatesMeshGlooSoloIov1Alpha2IssuedCertificates).AddMulticlusterIssuedCertificateReconciler(ctx, &remoteCertificatesMeshGlooSoloIoInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize CertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives reconcile loop for remote clusters
	certificates_mesh_gloo_solo_io_v1alpha2_controllers.NewMulticlusterPodBounceDirectiveReconcileLoop("PodBounceDirective", clusters, options.Remote.CertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives).AddMulticlusterPodBounceDirectiveReconciler(ctx, &remoteCertificatesMeshGlooSoloIoInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize XdsAgentEnterpriseMeshGlooSoloIov1Alpha1XdsConfigs reconcile loop for remote clusters
	xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1_controllers.NewMulticlusterXdsConfigReconcileLoop("XdsConfig", clusters, options.Remote.XdsAgentEnterpriseMeshGlooSoloIov1Alpha1XdsConfigs).AddMulticlusterXdsConfigReconciler(ctx, &remoteXdsAgentEnterpriseMeshGlooSoloIoInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize AppmeshK8SAwsv1Beta2VirtualServices reconcile loop for remote clusters
	appmesh_k8s_aws_v1beta2_controllers.NewMulticlusterVirtualServiceReconcileLoop("VirtualService", clusters, options.Remote.AppmeshK8SAwsv1Beta2VirtualServices).AddMulticlusterVirtualServiceReconciler(ctx, &remoteAppmeshK8SAwsInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize AppmeshK8SAwsv1Beta2VirtualNodes reconcile loop for remote clusters
	appmesh_k8s_aws_v1beta2_controllers.NewMulticlusterVirtualNodeReconcileLoop("VirtualNode", clusters, options.Remote.AppmeshK8SAwsv1Beta2VirtualNodes).AddMulticlusterVirtualNodeReconciler(ctx, &remoteAppmeshK8SAwsInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize AppmeshK8SAwsv1Beta2VirtualRouters reconcile loop for remote clusters
	appmesh_k8s_aws_v1beta2_controllers.NewMulticlusterVirtualRouterReconcileLoop("VirtualRouter", clusters, options.Remote.AppmeshK8SAwsv1Beta2VirtualRouters).AddMulticlusterVirtualRouterReconciler(ctx, &remoteAppmeshK8SAwsInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize NetworkingIstioIov1Alpha3DestinationRules reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterDestinationRuleReconcileLoop("DestinationRule", clusters, options.Remote.NetworkingIstioIov1Alpha3DestinationRules).AddMulticlusterDestinationRuleReconciler(ctx, &remoteNetworkingIstioIoInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize NetworkingIstioIov1Alpha3EnvoyFilters reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterEnvoyFilterReconcileLoop("EnvoyFilter", clusters, options.Remote.NetworkingIstioIov1Alpha3EnvoyFilters).AddMulticlusterEnvoyFilterReconciler(ctx, &remoteNetworkingIstioIoInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize NetworkingIstioIov1Alpha3Gateways reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterGatewayReconcileLoop("Gateway", clusters, options.Remote.NetworkingIstioIov1Alpha3Gateways).AddMulticlusterGatewayReconciler(ctx, &remoteNetworkingIstioIoInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize NetworkingIstioIov1Alpha3ServiceEntries reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterServiceEntryReconcileLoop("ServiceEntry", clusters, options.Remote.NetworkingIstioIov1Alpha3ServiceEntries).AddMulticlusterServiceEntryReconciler(ctx, &remoteNetworkingIstioIoInputReconciler{base: base}, options.Remote.Predicates...)
	// initialize NetworkingIstioIov1Alpha3VirtualServices reconcile loop for remote clusters
	networking_istio_io_v1alpha3_controllers.NewMulticlusterVirtualServiceReconcileLoop("VirtualService", clusters, options.Remote.NetworkingIstioIov1Alpha3VirtualServices).AddMulticlusterVirtualServiceReconciler(ctx, &remoteNetworkingIstioIoInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize SecurityIstioIov1Beta1AuthorizationPolicies reconcile loop for remote clusters
	security_istio_io_v1beta1_controllers.NewMulticlusterAuthorizationPolicyReconcileLoop("AuthorizationPolicy", clusters, options.Remote.SecurityIstioIov1Beta1AuthorizationPolicies).AddMulticlusterAuthorizationPolicyReconciler(ctx, &remoteSecurityIstioIoInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize V1ConfigMaps reconcile loop for remote clusters
	v1_controllers.NewMulticlusterConfigMapReconcileLoop("ConfigMap", clusters, options.Remote.V1ConfigMaps).AddMulticlusterConfigMapReconciler(ctx, &remoteInputReconciler{base: base}, options.Remote.Predicates...)

	// initialize SettingsMeshGlooSoloIov1Alpha2Settings reconcile loop for local cluster
	if err := settings_mesh_gloo_solo_io_v1alpha2_controllers.NewSettingsReconcileLoop("Settings", mgr, options.Local.SettingsMeshGlooSoloIov1Alpha2Settings).RunSettingsReconciler(ctx, &localSettingsMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize DiscoveryMeshGlooSoloIov1Alpha2TrafficTargets reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1alpha2_controllers.NewTrafficTargetReconcileLoop("TrafficTarget", mgr, options.Local.DiscoveryMeshGlooSoloIov1Alpha2TrafficTargets).RunTrafficTargetReconciler(ctx, &localDiscoveryMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize DiscoveryMeshGlooSoloIov1Alpha2Workloads reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1alpha2_controllers.NewWorkloadReconcileLoop("Workload", mgr, options.Local.DiscoveryMeshGlooSoloIov1Alpha2Workloads).RunWorkloadReconciler(ctx, &localDiscoveryMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize DiscoveryMeshGlooSoloIov1Alpha2Meshes reconcile loop for local cluster
	if err := discovery_mesh_gloo_solo_io_v1alpha2_controllers.NewMeshReconcileLoop("Mesh", mgr, options.Local.DiscoveryMeshGlooSoloIov1Alpha2Meshes).RunMeshReconciler(ctx, &localDiscoveryMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize NetworkingMeshGlooSoloIov1Alpha2TrafficPolicies reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1alpha2_controllers.NewTrafficPolicyReconcileLoop("TrafficPolicy", mgr, options.Local.NetworkingMeshGlooSoloIov1Alpha2TrafficPolicies).RunTrafficPolicyReconciler(ctx, &localNetworkingMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize NetworkingMeshGlooSoloIov1Alpha2AccessPolicies reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1alpha2_controllers.NewAccessPolicyReconcileLoop("AccessPolicy", mgr, options.Local.NetworkingMeshGlooSoloIov1Alpha2AccessPolicies).RunAccessPolicyReconciler(ctx, &localNetworkingMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize NetworkingMeshGlooSoloIov1Alpha2VirtualMeshes reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1alpha2_controllers.NewVirtualMeshReconcileLoop("VirtualMesh", mgr, options.Local.NetworkingMeshGlooSoloIov1Alpha2VirtualMeshes).RunVirtualMeshReconciler(ctx, &localNetworkingMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}
	// initialize NetworkingMeshGlooSoloIov1Alpha2FailoverServices reconcile loop for local cluster
	if err := networking_mesh_gloo_solo_io_v1alpha2_controllers.NewFailoverServiceReconcileLoop("FailoverService", mgr, options.Local.NetworkingMeshGlooSoloIov1Alpha2FailoverServices).RunFailoverServiceReconciler(ctx, &localNetworkingMeshGlooSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize V1Secrets reconcile loop for local cluster
	if err := v1_controllers.NewSecretReconcileLoop("Secret", mgr, options.Local.V1Secrets).RunSecretReconciler(ctx, &localInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	// initialize MulticlusterSoloIov1Alpha1KubernetesClusters reconcile loop for local cluster
	if err := multicluster_solo_io_v1alpha1_controllers.NewKubernetesClusterReconcileLoop("KubernetesCluster", mgr, options.Local.MulticlusterSoloIov1Alpha1KubernetesClusters).RunKubernetesClusterReconciler(ctx, &localMulticlusterSoloIoInputReconciler{base: base}, options.Local.Predicates...); err != nil {
		return nil, err
	}

	return base, nil
}

// Options for reconciling a snapshot in remote clusters
type RemoteReconcileOptions struct {

	// Options for reconciling CertificatesMeshGlooSoloIov1Alpha2IssuedCertificates
	CertificatesMeshGlooSoloIov1Alpha2IssuedCertificates reconcile.Options
	// Options for reconciling CertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives
	CertificatesMeshGlooSoloIov1Alpha2PodBounceDirectives reconcile.Options

	// Options for reconciling XdsAgentEnterpriseMeshGlooSoloIov1Alpha1XdsConfigs
	XdsAgentEnterpriseMeshGlooSoloIov1Alpha1XdsConfigs reconcile.Options

	// Options for reconciling AppmeshK8SAwsv1Beta2VirtualServices
	AppmeshK8SAwsv1Beta2VirtualServices reconcile.Options
	// Options for reconciling AppmeshK8SAwsv1Beta2VirtualNodes
	AppmeshK8SAwsv1Beta2VirtualNodes reconcile.Options
	// Options for reconciling AppmeshK8SAwsv1Beta2VirtualRouters
	AppmeshK8SAwsv1Beta2VirtualRouters reconcile.Options

	// Options for reconciling NetworkingIstioIov1Alpha3DestinationRules
	NetworkingIstioIov1Alpha3DestinationRules reconcile.Options
	// Options for reconciling NetworkingIstioIov1Alpha3EnvoyFilters
	NetworkingIstioIov1Alpha3EnvoyFilters reconcile.Options
	// Options for reconciling NetworkingIstioIov1Alpha3Gateways
	NetworkingIstioIov1Alpha3Gateways reconcile.Options
	// Options for reconciling NetworkingIstioIov1Alpha3ServiceEntries
	NetworkingIstioIov1Alpha3ServiceEntries reconcile.Options
	// Options for reconciling NetworkingIstioIov1Alpha3VirtualServices
	NetworkingIstioIov1Alpha3VirtualServices reconcile.Options

	// Options for reconciling SecurityIstioIov1Beta1AuthorizationPolicies
	SecurityIstioIov1Beta1AuthorizationPolicies reconcile.Options

	// Options for reconciling V1ConfigMaps
	V1ConfigMaps reconcile.Options

	// optional predicates for filtering remote events
	Predicates []predicate.Predicate
}

type remoteCertificatesMeshGlooSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteCertificatesMeshGlooSoloIoInputReconciler) ReconcileIssuedCertificate(clusterName string, obj *certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificate) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteCertificatesMeshGlooSoloIoInputReconciler) ReconcileIssuedCertificateDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteCertificatesMeshGlooSoloIoInputReconciler) ReconcilePodBounceDirective(clusterName string, obj *certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirective) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteCertificatesMeshGlooSoloIoInputReconciler) ReconcilePodBounceDirectiveDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

type remoteXdsAgentEnterpriseMeshGlooSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteXdsAgentEnterpriseMeshGlooSoloIoInputReconciler) ReconcileXdsConfig(clusterName string, obj *xds_agent_enterprise_mesh_gloo_solo_io_v1alpha1.XdsConfig) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteXdsAgentEnterpriseMeshGlooSoloIoInputReconciler) ReconcileXdsConfigDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

type remoteAppmeshK8SAwsInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualService(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualService) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualNode(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualNode) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualNodeDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualRouter(clusterName string, obj *appmesh_k8s_aws_v1beta2.VirtualRouter) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteAppmeshK8SAwsInputReconciler) ReconcileVirtualRouterDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

type remoteNetworkingIstioIoInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileDestinationRuleDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileEnvoyFilter(clusterName string, obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileEnvoyFilterDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileGateway(clusterName string, obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileGatewayDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileServiceEntryDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteNetworkingIstioIoInputReconciler) ReconcileVirtualServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

type remoteSecurityIstioIoInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteSecurityIstioIoInputReconciler) ReconcileAuthorizationPolicy(clusterName string, obj *security_istio_io_v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteSecurityIstioIoInputReconciler) ReconcileAuthorizationPolicyDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

type remoteInputReconciler struct {
	base input.InputReconciler
}

func (r *remoteInputReconciler) ReconcileConfigMap(clusterName string, obj *v1.ConfigMap) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileRemoteGeneric(obj)
}

func (r *remoteInputReconciler) ReconcileConfigMapDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileRemoteGeneric(ref)
	return err
}

// Options for reconciling a snapshot in remote clusters
type LocalReconcileOptions struct {

	// Options for reconciling SettingsMeshGlooSoloIov1Alpha2Settings
	SettingsMeshGlooSoloIov1Alpha2Settings reconcile.Options

	// Options for reconciling DiscoveryMeshGlooSoloIov1Alpha2TrafficTargets
	DiscoveryMeshGlooSoloIov1Alpha2TrafficTargets reconcile.Options
	// Options for reconciling DiscoveryMeshGlooSoloIov1Alpha2Workloads
	DiscoveryMeshGlooSoloIov1Alpha2Workloads reconcile.Options
	// Options for reconciling DiscoveryMeshGlooSoloIov1Alpha2Meshes
	DiscoveryMeshGlooSoloIov1Alpha2Meshes reconcile.Options

	// Options for reconciling NetworkingMeshGlooSoloIov1Alpha2TrafficPolicies
	NetworkingMeshGlooSoloIov1Alpha2TrafficPolicies reconcile.Options
	// Options for reconciling NetworkingMeshGlooSoloIov1Alpha2AccessPolicies
	NetworkingMeshGlooSoloIov1Alpha2AccessPolicies reconcile.Options
	// Options for reconciling NetworkingMeshGlooSoloIov1Alpha2VirtualMeshes
	NetworkingMeshGlooSoloIov1Alpha2VirtualMeshes reconcile.Options
	// Options for reconciling NetworkingMeshGlooSoloIov1Alpha2FailoverServices
	NetworkingMeshGlooSoloIov1Alpha2FailoverServices reconcile.Options

	// Options for reconciling V1Secrets
	V1Secrets reconcile.Options

	// Options for reconciling MulticlusterSoloIov1Alpha1KubernetesClusters
	MulticlusterSoloIov1Alpha1KubernetesClusters reconcile.Options

	// optional predicates for filtering local events
	Predicates []predicate.Predicate
}

type localSettingsMeshGlooSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *localSettingsMeshGlooSoloIoInputReconciler) ReconcileSettings(obj *settings_mesh_gloo_solo_io_v1alpha2.Settings) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localSettingsMeshGlooSoloIoInputReconciler) ReconcileSettingsDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

type localDiscoveryMeshGlooSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileTrafficTarget(obj *discovery_mesh_gloo_solo_io_v1alpha2.TrafficTarget) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileTrafficTargetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileWorkload(obj *discovery_mesh_gloo_solo_io_v1alpha2.Workload) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileWorkloadDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileMesh(obj *discovery_mesh_gloo_solo_io_v1alpha2.Mesh) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localDiscoveryMeshGlooSoloIoInputReconciler) ReconcileMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

type localNetworkingMeshGlooSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileTrafficPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileAccessPolicy(obj *networking_mesh_gloo_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileAccessPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileVirtualMesh(obj *networking_mesh_gloo_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileVirtualMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileFailoverService(obj *networking_mesh_gloo_solo_io_v1alpha2.FailoverService) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localNetworkingMeshGlooSoloIoInputReconciler) ReconcileFailoverServiceDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

type localInputReconciler struct {
	base input.InputReconciler
}

func (r *localInputReconciler) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localInputReconciler) ReconcileSecretDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}

type localMulticlusterSoloIoInputReconciler struct {
	base input.InputReconciler
}

func (r *localMulticlusterSoloIoInputReconciler) ReconcileKubernetesCluster(obj *multicluster_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	return r.base.ReconcileLocalGeneric(obj)
}

func (r *localMulticlusterSoloIoInputReconciler) ReconcileKubernetesClusterDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileLocalGeneric(ref)
	return err
}
