// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// The Input Snapshot contains the set of all:
// * Settings
// * TrafficTargets
// * Workloads
// * Meshes
// * TrafficPolicies
// * AccessPolicies
// * VirtualMeshes
// * FailoverServices
// * Secrets
// * KubernetesClusters
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	settings_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"
	settings_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"

	discovery_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	discovery_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"

	networking_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	networking_mesh_gloo_solo_io_v1alpha2_sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
)

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input Settings
	Settings() settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet

	// return the set of input TrafficTargets
	TrafficTargets() discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet
	// return the set of input Workloads
	Workloads() discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet
	// return the set of input Meshes
	Meshes() discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet

	// return the set of input TrafficPolicies
	TrafficPolicies() networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet
	// return the set of input AccessPolicies
	AccessPolicies() networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet
	// return the set of input VirtualMeshes
	VirtualMeshes() networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet
	// return the set of input FailoverServices
	FailoverServices() networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet

	// return the set of input Secrets
	Secrets() v1_sets.SecretSet

	// return the set of input KubernetesClusters
	KubernetesClusters() multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
	// update the status of all input objects which support
	// the Status subresource (in the local cluster)
	SyncStatuses(ctx context.Context, c client.Client, opts SyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type SyncStatusOptions struct {

	// sync status of Settings objects
	Settings bool

	// sync status of TrafficTarget objects
	TrafficTarget bool
	// sync status of Workload objects
	Workload bool
	// sync status of Mesh objects
	Mesh bool

	// sync status of TrafficPolicy objects
	TrafficPolicy bool
	// sync status of AccessPolicy objects
	AccessPolicy bool
	// sync status of VirtualMesh objects
	VirtualMesh bool
	// sync status of FailoverService objects
	FailoverService bool

	// sync status of Secret objects
	Secret bool

	// sync status of KubernetesCluster objects
	KubernetesCluster bool
}

type snapshot struct {
	name string

	settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet

	trafficTargets discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet
	workloads      discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet
	meshes         discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet

	trafficPolicies  networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet
	accessPolicies   networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet
	virtualMeshes    networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet
	failoverServices networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet

	secrets v1_sets.SecretSet

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet
}

func NewSnapshot(
	name string,

	settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet,

	trafficTargets discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet,
	workloads discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet,
	meshes discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet,

	trafficPolicies networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet,
	accessPolicies networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet,
	virtualMeshes networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet,
	failoverServices networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet,

	secrets v1_sets.SecretSet,

	kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet,

) Snapshot {
	return &snapshot{
		name: name,

		settings:           settings,
		trafficTargets:     trafficTargets,
		workloads:          workloads,
		meshes:             meshes,
		trafficPolicies:    trafficPolicies,
		accessPolicies:     accessPolicies,
		virtualMeshes:      virtualMeshes,
		failoverServices:   failoverServices,
		secrets:            secrets,
		kubernetesClusters: kubernetesClusters,
	}
}

func (s snapshot) Settings() settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet {
	return s.settings
}

func (s snapshot) TrafficTargets() discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet {
	return s.trafficTargets
}

func (s snapshot) Workloads() discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet {
	return s.workloads
}

func (s snapshot) Meshes() discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet {
	return s.meshes
}

func (s snapshot) TrafficPolicies() networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet {
	return s.trafficPolicies
}

func (s snapshot) AccessPolicies() networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet {
	return s.accessPolicies
}

func (s snapshot) VirtualMeshes() networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet {
	return s.virtualMeshes
}

func (s snapshot) FailoverServices() networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet {
	return s.failoverServices
}

func (s snapshot) Secrets() v1_sets.SecretSet {
	return s.secrets
}

func (s snapshot) KubernetesClusters() multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet {
	return s.kubernetesClusters
}

func (s snapshot) SyncStatuses(ctx context.Context, c client.Client, opts SyncStatusOptions) error {
	var errs error

	if opts.Settings {
		for _, obj := range s.Settings().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.TrafficTarget {
		for _, obj := range s.TrafficTargets().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.Workload {
		for _, obj := range s.Workloads().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.Mesh {
		for _, obj := range s.Meshes().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.TrafficPolicy {
		for _, obj := range s.TrafficPolicies().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.AccessPolicy {
		for _, obj := range s.AccessPolicies().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.VirtualMesh {
		for _, obj := range s.VirtualMeshes().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.FailoverService {
		for _, obj := range s.FailoverServices().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.KubernetesCluster {
		for _, obj := range s.KubernetesClusters().List() {
			if _, err := controllerutils.UpdateStatus(ctx, c, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["settings"] = s.settings.List()
	snapshotMap["trafficTargets"] = s.trafficTargets.List()
	snapshotMap["workloads"] = s.workloads.List()
	snapshotMap["meshes"] = s.meshes.List()
	snapshotMap["trafficPolicies"] = s.trafficPolicies.List()
	snapshotMap["accessPolicies"] = s.accessPolicies.List()
	snapshotMap["virtualMeshes"] = s.virtualMeshes.List()
	snapshotMap["failoverServices"] = s.failoverServices.List()
	snapshotMap["secrets"] = s.secrets.List()
	snapshotMap["kubernetesClusters"] = s.kubernetesClusters.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type Builder interface {
	BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error)
}

// Options for building a snapshot
type BuildOptions struct {

	// List options for composing a snapshot from Settings
	Settings ResourceBuildOptions

	// List options for composing a snapshot from TrafficTargets
	TrafficTargets ResourceBuildOptions
	// List options for composing a snapshot from Workloads
	Workloads ResourceBuildOptions
	// List options for composing a snapshot from Meshes
	Meshes ResourceBuildOptions

	// List options for composing a snapshot from TrafficPolicies
	TrafficPolicies ResourceBuildOptions
	// List options for composing a snapshot from AccessPolicies
	AccessPolicies ResourceBuildOptions
	// List options for composing a snapshot from VirtualMeshes
	VirtualMeshes ResourceBuildOptions
	// List options for composing a snapshot from FailoverServices
	FailoverServices ResourceBuildOptions

	// List options for composing a snapshot from Secrets
	Secrets ResourceBuildOptions

	// List options for composing a snapshot from KubernetesClusters
	KubernetesClusters ResourceBuildOptions
}

// Options for reading resources of a given type
type ResourceBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources in a single cluster
type singleClusterBuilder struct {
	mgr manager.Manager
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewSingleClusterBuilder(
	mgr manager.Manager,
) Builder {
	return &singleClusterBuilder{
		mgr: mgr,
	}
}

func (b *singleClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	settings := settings_mesh_gloo_solo_io_v1alpha2_sets.NewSettingsSet()

	trafficTargets := discovery_mesh_gloo_solo_io_v1alpha2_sets.NewTrafficTargetSet()
	workloads := discovery_mesh_gloo_solo_io_v1alpha2_sets.NewWorkloadSet()
	meshes := discovery_mesh_gloo_solo_io_v1alpha2_sets.NewMeshSet()

	trafficPolicies := networking_mesh_gloo_solo_io_v1alpha2_sets.NewTrafficPolicySet()
	accessPolicies := networking_mesh_gloo_solo_io_v1alpha2_sets.NewAccessPolicySet()
	virtualMeshes := networking_mesh_gloo_solo_io_v1alpha2_sets.NewVirtualMeshSet()
	failoverServices := networking_mesh_gloo_solo_io_v1alpha2_sets.NewFailoverServiceSet()

	secrets := v1_sets.NewSecretSet()

	kubernetesClusters := multicluster_solo_io_v1alpha1_sets.NewKubernetesClusterSet()

	var errs error

	if err := b.insertSettings(ctx, settings, opts.Settings); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertTrafficTargets(ctx, trafficTargets, opts.TrafficTargets); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertWorkloads(ctx, workloads, opts.Workloads); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertMeshes(ctx, meshes, opts.Meshes); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertTrafficPolicies(ctx, trafficPolicies, opts.TrafficPolicies); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertAccessPolicies(ctx, accessPolicies, opts.AccessPolicies); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertVirtualMeshes(ctx, virtualMeshes, opts.VirtualMeshes); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertFailoverServices(ctx, failoverServices, opts.FailoverServices); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertSecrets(ctx, secrets, opts.Secrets); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertKubernetesClusters(ctx, kubernetesClusters, opts.KubernetesClusters); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSnapshot(
		name,

		settings,
		trafficTargets,
		workloads,
		meshes,
		trafficPolicies,
		accessPolicies,
		virtualMeshes,
		failoverServices,
		secrets,
		kubernetesClusters,
	)

	return outputSnap, errs
}

func (b *singleClusterBuilder) insertSettings(ctx context.Context, settings settings_mesh_gloo_solo_io_v1alpha2_sets.SettingsSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "settings.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "Settings",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	settingsList, err := settings_mesh_gloo_solo_io_v1alpha2.NewSettingsClient(b.mgr.GetClient()).ListSettings(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range settingsList.Items {
		item := item // pike
		settings.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertTrafficTargets(ctx context.Context, trafficTargets discovery_mesh_gloo_solo_io_v1alpha2_sets.TrafficTargetSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "TrafficTarget",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	trafficTargetList, err := discovery_mesh_gloo_solo_io_v1alpha2.NewTrafficTargetClient(b.mgr.GetClient()).ListTrafficTarget(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range trafficTargetList.Items {
		item := item // pike
		trafficTargets.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertWorkloads(ctx context.Context, workloads discovery_mesh_gloo_solo_io_v1alpha2_sets.WorkloadSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "Workload",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	workloadList, err := discovery_mesh_gloo_solo_io_v1alpha2.NewWorkloadClient(b.mgr.GetClient()).ListWorkload(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range workloadList.Items {
		item := item // pike
		workloads.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertMeshes(ctx context.Context, meshes discovery_mesh_gloo_solo_io_v1alpha2_sets.MeshSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "discovery.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "Mesh",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	meshList, err := discovery_mesh_gloo_solo_io_v1alpha2.NewMeshClient(b.mgr.GetClient()).ListMesh(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range meshList.Items {
		item := item // pike
		meshes.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertTrafficPolicies(ctx context.Context, trafficPolicies networking_mesh_gloo_solo_io_v1alpha2_sets.TrafficPolicySet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "networking.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "TrafficPolicy",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	trafficPolicyList, err := networking_mesh_gloo_solo_io_v1alpha2.NewTrafficPolicyClient(b.mgr.GetClient()).ListTrafficPolicy(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range trafficPolicyList.Items {
		item := item // pike
		trafficPolicies.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertAccessPolicies(ctx context.Context, accessPolicies networking_mesh_gloo_solo_io_v1alpha2_sets.AccessPolicySet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "networking.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "AccessPolicy",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	accessPolicyList, err := networking_mesh_gloo_solo_io_v1alpha2.NewAccessPolicyClient(b.mgr.GetClient()).ListAccessPolicy(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range accessPolicyList.Items {
		item := item // pike
		accessPolicies.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertVirtualMeshes(ctx context.Context, virtualMeshes networking_mesh_gloo_solo_io_v1alpha2_sets.VirtualMeshSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "networking.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "VirtualMesh",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	virtualMeshList, err := networking_mesh_gloo_solo_io_v1alpha2.NewVirtualMeshClient(b.mgr.GetClient()).ListVirtualMesh(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range virtualMeshList.Items {
		item := item // pike
		virtualMeshes.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertFailoverServices(ctx context.Context, failoverServices networking_mesh_gloo_solo_io_v1alpha2_sets.FailoverServiceSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "networking.mesh.gloo.solo.io",
			Version: "v1alpha2",
			Kind:    "FailoverService",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	failoverServiceList, err := networking_mesh_gloo_solo_io_v1alpha2.NewFailoverServiceClient(b.mgr.GetClient()).ListFailoverService(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range failoverServiceList.Items {
		item := item // pike
		failoverServices.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertSecrets(ctx context.Context, secrets v1_sets.SecretSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Secret",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	secretList, err := v1.NewSecretClient(b.mgr.GetClient()).ListSecret(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range secretList.Items {
		item := item // pike
		secrets.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertKubernetesClusters(ctx context.Context, kubernetesClusters multicluster_solo_io_v1alpha1_sets.KubernetesClusterSet, opts ResourceBuildOptions) error {

	if opts.Verifier != nil {
		gvk := schema.GroupVersionKind{
			Group:   "multicluster.solo.io",
			Version: "v1alpha1",
			Kind:    "KubernetesCluster",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			"", // verify in the local cluster
			b.mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	kubernetesClusterList, err := multicluster_solo_io_v1alpha1.NewKubernetesClusterClient(b.mgr.GetClient()).ListKubernetesCluster(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range kubernetesClusterList.Items {
		item := item // pike
		kubernetesClusters.Insert(&item)
	}

	return nil
}