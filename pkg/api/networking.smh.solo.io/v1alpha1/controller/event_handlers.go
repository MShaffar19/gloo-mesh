// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the TrafficPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type TrafficPolicyEventHandler interface {
	CreateTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	UpdateTrafficPolicy(old, new *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	DeleteTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	GenericTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
}

type TrafficPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	OnUpdate  func(old, new *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	OnDelete  func(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
	OnGeneric func(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error
}

func (f *TrafficPolicyEventHandlerFuncs) CreateTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *TrafficPolicyEventHandlerFuncs) DeleteTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *TrafficPolicyEventHandlerFuncs) UpdateTrafficPolicy(objOld, objNew *networking_smh_solo_io_v1alpha1.TrafficPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *TrafficPolicyEventHandlerFuncs) GenericTrafficPolicy(obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type TrafficPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h TrafficPolicyEventHandler, predicates ...predicate.Predicate) error
}

type trafficPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewTrafficPolicyEventWatcher(name string, mgr manager.Manager) TrafficPolicyEventWatcher {
	return &trafficPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_smh_solo_io_v1alpha1.TrafficPolicy{}),
	}
}

func (c *trafficPolicyEventWatcher) AddEventHandler(ctx context.Context, h TrafficPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTrafficPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericTrafficPolicyHandler implements a generic events.EventHandler
type genericTrafficPolicyHandler struct {
	handler TrafficPolicyEventHandler
}

func (h genericTrafficPolicyHandler) Create(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return h.handler.CreateTrafficPolicy(obj)
}

func (h genericTrafficPolicyHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return h.handler.DeleteTrafficPolicy(obj)
}

func (h genericTrafficPolicyHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", new)
	}
	return h.handler.UpdateTrafficPolicy(objOld, objNew)
}

func (h genericTrafficPolicyHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return h.handler.GenericTrafficPolicy(obj)
}

// Handle events for the AccessControlPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type AccessControlPolicyEventHandler interface {
	CreateAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	UpdateAccessControlPolicy(old, new *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	DeleteAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	GenericAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
}

type AccessControlPolicyEventHandlerFuncs struct {
	OnCreate  func(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	OnUpdate  func(old, new *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	OnDelete  func(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
	OnGeneric func(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error
}

func (f *AccessControlPolicyEventHandlerFuncs) CreateAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *AccessControlPolicyEventHandlerFuncs) DeleteAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *AccessControlPolicyEventHandlerFuncs) UpdateAccessControlPolicy(objOld, objNew *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *AccessControlPolicyEventHandlerFuncs) GenericAccessControlPolicy(obj *networking_smh_solo_io_v1alpha1.AccessControlPolicy) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type AccessControlPolicyEventWatcher interface {
	AddEventHandler(ctx context.Context, h AccessControlPolicyEventHandler, predicates ...predicate.Predicate) error
}

type accessControlPolicyEventWatcher struct {
	watcher events.EventWatcher
}

func NewAccessControlPolicyEventWatcher(name string, mgr manager.Manager) AccessControlPolicyEventWatcher {
	return &accessControlPolicyEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_smh_solo_io_v1alpha1.AccessControlPolicy{}),
	}
}

func (c *accessControlPolicyEventWatcher) AddEventHandler(ctx context.Context, h AccessControlPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericAccessControlPolicyHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericAccessControlPolicyHandler implements a generic events.EventHandler
type genericAccessControlPolicyHandler struct {
	handler AccessControlPolicyEventHandler
}

func (h genericAccessControlPolicyHandler) Create(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return h.handler.CreateAccessControlPolicy(obj)
}

func (h genericAccessControlPolicyHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return h.handler.DeleteAccessControlPolicy(obj)
}

func (h genericAccessControlPolicyHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", old)
	}
	objNew, ok := new.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", new)
	}
	return h.handler.UpdateAccessControlPolicy(objOld, objNew)
}

func (h genericAccessControlPolicyHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return h.handler.GenericAccessControlPolicy(obj)
}

// Handle events for the VirtualMesh Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualMeshEventHandler interface {
	CreateVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	UpdateVirtualMesh(old, new *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	DeleteVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	GenericVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
}

type VirtualMeshEventHandlerFuncs struct {
	OnCreate  func(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	OnUpdate  func(old, new *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	OnDelete  func(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
	OnGeneric func(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error
}

func (f *VirtualMeshEventHandlerFuncs) CreateVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *VirtualMeshEventHandlerFuncs) DeleteVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *VirtualMeshEventHandlerFuncs) UpdateVirtualMesh(objOld, objNew *networking_smh_solo_io_v1alpha1.VirtualMesh) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *VirtualMeshEventHandlerFuncs) GenericVirtualMesh(obj *networking_smh_solo_io_v1alpha1.VirtualMesh) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type VirtualMeshEventWatcher interface {
	AddEventHandler(ctx context.Context, h VirtualMeshEventHandler, predicates ...predicate.Predicate) error
}

type virtualMeshEventWatcher struct {
	watcher events.EventWatcher
}

func NewVirtualMeshEventWatcher(name string, mgr manager.Manager) VirtualMeshEventWatcher {
	return &virtualMeshEventWatcher{
		watcher: events.NewWatcher(name, mgr, &networking_smh_solo_io_v1alpha1.VirtualMesh{}),
	}
}

func (c *virtualMeshEventWatcher) AddEventHandler(ctx context.Context, h VirtualMeshEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualMeshHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericVirtualMeshHandler implements a generic events.EventHandler
type genericVirtualMeshHandler struct {
	handler VirtualMeshEventHandler
}

func (h genericVirtualMeshHandler) Create(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return h.handler.CreateVirtualMesh(obj)
}

func (h genericVirtualMeshHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return h.handler.DeleteVirtualMesh(obj)
}

func (h genericVirtualMeshHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", old)
	}
	objNew, ok := new.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", new)
	}
	return h.handler.UpdateVirtualMesh(objOld, objNew)
}

func (h genericVirtualMeshHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return h.handler.GenericVirtualMesh(obj)
}