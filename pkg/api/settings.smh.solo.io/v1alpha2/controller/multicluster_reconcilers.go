// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	settings_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/settings.smh.solo.io/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Settings Resource across clusters.
// implemented by the user
type MulticlusterSettingsReconciler interface {
	ReconcileSettings(clusterName string, obj *settings_smh_solo_io_v1alpha2.Settings) (reconcile.Result, error)
}

// Reconcile deletion events for the Settings Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterSettingsDeletionReconciler interface {
	ReconcileSettingsDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterSettingsReconcilerFuncs struct {
	OnReconcileSettings         func(clusterName string, obj *settings_smh_solo_io_v1alpha2.Settings) (reconcile.Result, error)
	OnReconcileSettingsDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterSettingsReconcilerFuncs) ReconcileSettings(clusterName string, obj *settings_smh_solo_io_v1alpha2.Settings) (reconcile.Result, error) {
	if f.OnReconcileSettings == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSettings(clusterName, obj)
}

func (f *MulticlusterSettingsReconcilerFuncs) ReconcileSettingsDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileSettingsDeletion == nil {
		return nil
	}
	return f.OnReconcileSettingsDeletion(clusterName, req)
}

type MulticlusterSettingsReconcileLoop interface {
	// AddMulticlusterSettingsReconciler adds a MulticlusterSettingsReconciler to the MulticlusterSettingsReconcileLoop.
	AddMulticlusterSettingsReconciler(ctx context.Context, rec MulticlusterSettingsReconciler, predicates ...predicate.Predicate)
}

type multiclusterSettingsReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterSettingsReconcileLoop) AddMulticlusterSettingsReconciler(ctx context.Context, rec MulticlusterSettingsReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericSettingsMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterSettingsReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterSettingsReconcileLoop {
	return &multiclusterSettingsReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &settings_smh_solo_io_v1alpha2.Settings{}, options)}
}

type genericSettingsMulticlusterReconciler struct {
	reconciler MulticlusterSettingsReconciler
}

func (g genericSettingsMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterSettingsDeletionReconciler); ok {
		return deletionReconciler.ReconcileSettingsDeletion(cluster, req)
	}
	return nil
}

func (g genericSettingsMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*settings_smh_solo_io_v1alpha2.Settings)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Settings handler received event for %T", object)
	}
	return g.reconciler.ReconcileSettings(cluster, obj)
}
