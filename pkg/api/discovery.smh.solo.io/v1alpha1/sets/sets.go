// Code generated by skv2. DO NOT EDIT.

package v1alpha1sets

import (
	discovery_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type KubernetesClusterSet interface {
	Keys() sets.String
	List() []*discovery_smh_solo_io_v1alpha1.KubernetesCluster
	Map() map[string]*discovery_smh_solo_io_v1alpha1.KubernetesCluster
	Insert(kubernetesCluster ...*discovery_smh_solo_io_v1alpha1.KubernetesCluster)
	Equal(kubernetesClusterSet KubernetesClusterSet) bool
	Has(kubernetesCluster *discovery_smh_solo_io_v1alpha1.KubernetesCluster) bool
	Delete(kubernetesCluster *discovery_smh_solo_io_v1alpha1.KubernetesCluster)
	Union(set KubernetesClusterSet) KubernetesClusterSet
	Difference(set KubernetesClusterSet) KubernetesClusterSet
	Intersection(set KubernetesClusterSet) KubernetesClusterSet
}

func makeGenericKubernetesClusterSet(kubernetesClusterList []*discovery_smh_solo_io_v1alpha1.KubernetesCluster) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range kubernetesClusterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type kubernetesClusterSet struct {
	set sksets.ResourceSet
}

func NewKubernetesClusterSet(kubernetesClusterList ...*discovery_smh_solo_io_v1alpha1.KubernetesCluster) KubernetesClusterSet {
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(kubernetesClusterList)}
}

func (s kubernetesClusterSet) Keys() sets.String {
	return s.set.Keys()
}

func (s kubernetesClusterSet) List() []*discovery_smh_solo_io_v1alpha1.KubernetesCluster {
	var kubernetesClusterList []*discovery_smh_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range s.set.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*discovery_smh_solo_io_v1alpha1.KubernetesCluster))
	}
	return kubernetesClusterList
}

func (s kubernetesClusterSet) Map() map[string]*discovery_smh_solo_io_v1alpha1.KubernetesCluster {
	newMap := map[string]*discovery_smh_solo_io_v1alpha1.KubernetesCluster{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*discovery_smh_solo_io_v1alpha1.KubernetesCluster)
	}
	return newMap
}

func (s kubernetesClusterSet) Insert(
	kubernetesClusterList ...*discovery_smh_solo_io_v1alpha1.KubernetesCluster,
) {
	for _, obj := range kubernetesClusterList {
		s.set.Insert(obj)
	}
}

func (s kubernetesClusterSet) Has(kubernetesCluster *discovery_smh_solo_io_v1alpha1.KubernetesCluster) bool {
	return s.set.Has(kubernetesCluster)
}

func (s kubernetesClusterSet) Equal(
	kubernetesClusterSet KubernetesClusterSet,
) bool {
	return s.set.Equal(makeGenericKubernetesClusterSet(kubernetesClusterSet.List()))
}

func (s kubernetesClusterSet) Delete(KubernetesCluster *discovery_smh_solo_io_v1alpha1.KubernetesCluster) {
	s.set.Delete(KubernetesCluster)
}

func (s kubernetesClusterSet) Union(set KubernetesClusterSet) KubernetesClusterSet {
	return NewKubernetesClusterSet(append(s.List(), set.List()...)...)
}

func (s kubernetesClusterSet) Difference(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Difference(makeGenericKubernetesClusterSet(set.List()))
	return kubernetesClusterSet{set: newSet}
}

func (s kubernetesClusterSet) Intersection(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Intersection(makeGenericKubernetesClusterSet(set.List()))
	var kubernetesClusterList []*discovery_smh_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range newSet.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*discovery_smh_solo_io_v1alpha1.KubernetesCluster))
	}
	return NewKubernetesClusterSet(kubernetesClusterList...)
}

type MeshServiceSet interface {
	Keys() sets.String
	List() []*discovery_smh_solo_io_v1alpha1.MeshService
	Map() map[string]*discovery_smh_solo_io_v1alpha1.MeshService
	Insert(meshService ...*discovery_smh_solo_io_v1alpha1.MeshService)
	Equal(meshServiceSet MeshServiceSet) bool
	Has(meshService *discovery_smh_solo_io_v1alpha1.MeshService) bool
	Delete(meshService *discovery_smh_solo_io_v1alpha1.MeshService)
	Union(set MeshServiceSet) MeshServiceSet
	Difference(set MeshServiceSet) MeshServiceSet
	Intersection(set MeshServiceSet) MeshServiceSet
}

func makeGenericMeshServiceSet(meshServiceList []*discovery_smh_solo_io_v1alpha1.MeshService) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshServiceSet struct {
	set sksets.ResourceSet
}

func NewMeshServiceSet(meshServiceList ...*discovery_smh_solo_io_v1alpha1.MeshService) MeshServiceSet {
	return &meshServiceSet{set: makeGenericMeshServiceSet(meshServiceList)}
}

func (s meshServiceSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshServiceSet) List() []*discovery_smh_solo_io_v1alpha1.MeshService {
	var meshServiceList []*discovery_smh_solo_io_v1alpha1.MeshService
	for _, obj := range s.set.List() {
		meshServiceList = append(meshServiceList, obj.(*discovery_smh_solo_io_v1alpha1.MeshService))
	}
	return meshServiceList
}

func (s meshServiceSet) Map() map[string]*discovery_smh_solo_io_v1alpha1.MeshService {
	newMap := map[string]*discovery_smh_solo_io_v1alpha1.MeshService{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*discovery_smh_solo_io_v1alpha1.MeshService)
	}
	return newMap
}

func (s meshServiceSet) Insert(
	meshServiceList ...*discovery_smh_solo_io_v1alpha1.MeshService,
) {
	for _, obj := range meshServiceList {
		s.set.Insert(obj)
	}
}

func (s meshServiceSet) Has(meshService *discovery_smh_solo_io_v1alpha1.MeshService) bool {
	return s.set.Has(meshService)
}

func (s meshServiceSet) Equal(
	meshServiceSet MeshServiceSet,
) bool {
	return s.set.Equal(makeGenericMeshServiceSet(meshServiceSet.List()))
}

func (s meshServiceSet) Delete(MeshService *discovery_smh_solo_io_v1alpha1.MeshService) {
	s.set.Delete(MeshService)
}

func (s meshServiceSet) Union(set MeshServiceSet) MeshServiceSet {
	return NewMeshServiceSet(append(s.List(), set.List()...)...)
}

func (s meshServiceSet) Difference(set MeshServiceSet) MeshServiceSet {
	newSet := s.set.Difference(makeGenericMeshServiceSet(set.List()))
	return meshServiceSet{set: newSet}
}

func (s meshServiceSet) Intersection(set MeshServiceSet) MeshServiceSet {
	newSet := s.set.Intersection(makeGenericMeshServiceSet(set.List()))
	var meshServiceList []*discovery_smh_solo_io_v1alpha1.MeshService
	for _, obj := range newSet.List() {
		meshServiceList = append(meshServiceList, obj.(*discovery_smh_solo_io_v1alpha1.MeshService))
	}
	return NewMeshServiceSet(meshServiceList...)
}

type MeshWorkloadSet interface {
	Keys() sets.String
	List() []*discovery_smh_solo_io_v1alpha1.MeshWorkload
	Map() map[string]*discovery_smh_solo_io_v1alpha1.MeshWorkload
	Insert(meshWorkload ...*discovery_smh_solo_io_v1alpha1.MeshWorkload)
	Equal(meshWorkloadSet MeshWorkloadSet) bool
	Has(meshWorkload *discovery_smh_solo_io_v1alpha1.MeshWorkload) bool
	Delete(meshWorkload *discovery_smh_solo_io_v1alpha1.MeshWorkload)
	Union(set MeshWorkloadSet) MeshWorkloadSet
	Difference(set MeshWorkloadSet) MeshWorkloadSet
	Intersection(set MeshWorkloadSet) MeshWorkloadSet
}

func makeGenericMeshWorkloadSet(meshWorkloadList []*discovery_smh_solo_io_v1alpha1.MeshWorkload) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshWorkloadList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshWorkloadSet struct {
	set sksets.ResourceSet
}

func NewMeshWorkloadSet(meshWorkloadList ...*discovery_smh_solo_io_v1alpha1.MeshWorkload) MeshWorkloadSet {
	return &meshWorkloadSet{set: makeGenericMeshWorkloadSet(meshWorkloadList)}
}

func (s meshWorkloadSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshWorkloadSet) List() []*discovery_smh_solo_io_v1alpha1.MeshWorkload {
	var meshWorkloadList []*discovery_smh_solo_io_v1alpha1.MeshWorkload
	for _, obj := range s.set.List() {
		meshWorkloadList = append(meshWorkloadList, obj.(*discovery_smh_solo_io_v1alpha1.MeshWorkload))
	}
	return meshWorkloadList
}

func (s meshWorkloadSet) Map() map[string]*discovery_smh_solo_io_v1alpha1.MeshWorkload {
	newMap := map[string]*discovery_smh_solo_io_v1alpha1.MeshWorkload{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*discovery_smh_solo_io_v1alpha1.MeshWorkload)
	}
	return newMap
}

func (s meshWorkloadSet) Insert(
	meshWorkloadList ...*discovery_smh_solo_io_v1alpha1.MeshWorkload,
) {
	for _, obj := range meshWorkloadList {
		s.set.Insert(obj)
	}
}

func (s meshWorkloadSet) Has(meshWorkload *discovery_smh_solo_io_v1alpha1.MeshWorkload) bool {
	return s.set.Has(meshWorkload)
}

func (s meshWorkloadSet) Equal(
	meshWorkloadSet MeshWorkloadSet,
) bool {
	return s.set.Equal(makeGenericMeshWorkloadSet(meshWorkloadSet.List()))
}

func (s meshWorkloadSet) Delete(MeshWorkload *discovery_smh_solo_io_v1alpha1.MeshWorkload) {
	s.set.Delete(MeshWorkload)
}

func (s meshWorkloadSet) Union(set MeshWorkloadSet) MeshWorkloadSet {
	return NewMeshWorkloadSet(append(s.List(), set.List()...)...)
}

func (s meshWorkloadSet) Difference(set MeshWorkloadSet) MeshWorkloadSet {
	newSet := s.set.Difference(makeGenericMeshWorkloadSet(set.List()))
	return meshWorkloadSet{set: newSet}
}

func (s meshWorkloadSet) Intersection(set MeshWorkloadSet) MeshWorkloadSet {
	newSet := s.set.Intersection(makeGenericMeshWorkloadSet(set.List()))
	var meshWorkloadList []*discovery_smh_solo_io_v1alpha1.MeshWorkload
	for _, obj := range newSet.List() {
		meshWorkloadList = append(meshWorkloadList, obj.(*discovery_smh_solo_io_v1alpha1.MeshWorkload))
	}
	return NewMeshWorkloadSet(meshWorkloadList...)
}

type MeshSet interface {
	Keys() sets.String
	List() []*discovery_smh_solo_io_v1alpha1.Mesh
	Map() map[string]*discovery_smh_solo_io_v1alpha1.Mesh
	Insert(mesh ...*discovery_smh_solo_io_v1alpha1.Mesh)
	Equal(meshSet MeshSet) bool
	Has(mesh *discovery_smh_solo_io_v1alpha1.Mesh) bool
	Delete(mesh *discovery_smh_solo_io_v1alpha1.Mesh)
	Union(set MeshSet) MeshSet
	Difference(set MeshSet) MeshSet
	Intersection(set MeshSet) MeshSet
}

func makeGenericMeshSet(meshList []*discovery_smh_solo_io_v1alpha1.Mesh) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshSet struct {
	set sksets.ResourceSet
}

func NewMeshSet(meshList ...*discovery_smh_solo_io_v1alpha1.Mesh) MeshSet {
	return &meshSet{set: makeGenericMeshSet(meshList)}
}

func (s meshSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshSet) List() []*discovery_smh_solo_io_v1alpha1.Mesh {
	var meshList []*discovery_smh_solo_io_v1alpha1.Mesh
	for _, obj := range s.set.List() {
		meshList = append(meshList, obj.(*discovery_smh_solo_io_v1alpha1.Mesh))
	}
	return meshList
}

func (s meshSet) Map() map[string]*discovery_smh_solo_io_v1alpha1.Mesh {
	newMap := map[string]*discovery_smh_solo_io_v1alpha1.Mesh{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*discovery_smh_solo_io_v1alpha1.Mesh)
	}
	return newMap
}

func (s meshSet) Insert(
	meshList ...*discovery_smh_solo_io_v1alpha1.Mesh,
) {
	for _, obj := range meshList {
		s.set.Insert(obj)
	}
}

func (s meshSet) Has(mesh *discovery_smh_solo_io_v1alpha1.Mesh) bool {
	return s.set.Has(mesh)
}

func (s meshSet) Equal(
	meshSet MeshSet,
) bool {
	return s.set.Equal(makeGenericMeshSet(meshSet.List()))
}

func (s meshSet) Delete(Mesh *discovery_smh_solo_io_v1alpha1.Mesh) {
	s.set.Delete(Mesh)
}

func (s meshSet) Union(set MeshSet) MeshSet {
	return NewMeshSet(append(s.List(), set.List()...)...)
}

func (s meshSet) Difference(set MeshSet) MeshSet {
	newSet := s.set.Difference(makeGenericMeshSet(set.List()))
	return meshSet{set: newSet}
}

func (s meshSet) Intersection(set MeshSet) MeshSet {
	newSet := s.set.Intersection(makeGenericMeshSet(set.List()))
	var meshList []*discovery_smh_solo_io_v1alpha1.Mesh
	for _, obj := range newSet.List() {
		meshList = append(meshList, obj.(*discovery_smh_solo_io_v1alpha1.Mesh))
	}
	return NewMeshSet(meshList...)
}
