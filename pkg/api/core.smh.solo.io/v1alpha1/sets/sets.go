// Code generated by skv2. DO NOT EDIT.

package v1alpha1sets

import (
	core_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type SettingsSet interface {
	Keys() sets.String
	List() []*core_smh_solo_io_v1alpha1.Settings
	Map() map[string]*core_smh_solo_io_v1alpha1.Settings
	Insert(settings ...*core_smh_solo_io_v1alpha1.Settings)
	Equal(settingsSet SettingsSet) bool
	Has(settings *core_smh_solo_io_v1alpha1.Settings) bool
	Delete(settings *core_smh_solo_io_v1alpha1.Settings)
	Union(set SettingsSet) SettingsSet
	Difference(set SettingsSet) SettingsSet
	Intersection(set SettingsSet) SettingsSet
}

func makeGenericSettingsSet(settingsList []*core_smh_solo_io_v1alpha1.Settings) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range settingsList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type settingsSet struct {
	set sksets.ResourceSet
}

func NewSettingsSet(settingsList ...*core_smh_solo_io_v1alpha1.Settings) SettingsSet {
	return &settingsSet{set: makeGenericSettingsSet(settingsList)}
}

func (s settingsSet) Keys() sets.String {
	return s.set.Keys()
}

func (s settingsSet) List() []*core_smh_solo_io_v1alpha1.Settings {
	var settingsList []*core_smh_solo_io_v1alpha1.Settings
	for _, obj := range s.set.List() {
		settingsList = append(settingsList, obj.(*core_smh_solo_io_v1alpha1.Settings))
	}
	return settingsList
}

func (s settingsSet) Map() map[string]*core_smh_solo_io_v1alpha1.Settings {
	newMap := map[string]*core_smh_solo_io_v1alpha1.Settings{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*core_smh_solo_io_v1alpha1.Settings)
	}
	return newMap
}

func (s settingsSet) Insert(
	settingsList ...*core_smh_solo_io_v1alpha1.Settings,
) {
	for _, obj := range settingsList {
		s.set.Insert(obj)
	}
}

func (s settingsSet) Has(settings *core_smh_solo_io_v1alpha1.Settings) bool {
	return s.set.Has(settings)
}

func (s settingsSet) Equal(
	settingsSet SettingsSet,
) bool {
	return s.set.Equal(makeGenericSettingsSet(settingsSet.List()))
}

func (s settingsSet) Delete(Settings *core_smh_solo_io_v1alpha1.Settings) {
	s.set.Delete(Settings)
}

func (s settingsSet) Union(set SettingsSet) SettingsSet {
	return NewSettingsSet(append(s.List(), set.List()...)...)
}

func (s settingsSet) Difference(set SettingsSet) SettingsSet {
	newSet := s.set.Difference(makeGenericSettingsSet(set.List()))
	return settingsSet{set: newSet}
}

func (s settingsSet) Intersection(set SettingsSet) SettingsSet {
	newSet := s.set.Intersection(makeGenericSettingsSet(set.List()))
	var settingsList []*core_smh_solo_io_v1alpha1.Settings
	for _, obj := range newSet.List() {
		settingsList = append(settingsList, obj.(*core_smh_solo_io_v1alpha1.Settings))
	}
	return NewSettingsSet(settingsList...)
}
