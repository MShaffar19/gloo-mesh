// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the security.smh.solo.io/v1alpha1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the security.smh.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the security.smh.solo.io/v1alpha1/v1alpha1 APIs
	VirtualMeshCertificateSigningRequests() VirtualMeshCertificateSigningRequestClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the security.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) VirtualMeshCertificateSigningRequests() VirtualMeshCertificateSigningRequestClient {
	return NewVirtualMeshCertificateSigningRequestClient(c.client)
}

// Reader knows how to read and list VirtualMeshCertificateSigningRequests.
type VirtualMeshCertificateSigningRequestReader interface {
	// Get retrieves a VirtualMeshCertificateSigningRequest for the given object key
	GetVirtualMeshCertificateSigningRequest(ctx context.Context, key client.ObjectKey) (*VirtualMeshCertificateSigningRequest, error)

	// List retrieves list of VirtualMeshCertificateSigningRequests for a given namespace and list options.
	ListVirtualMeshCertificateSigningRequest(ctx context.Context, opts ...client.ListOption) (*VirtualMeshCertificateSigningRequestList, error)
}

// VirtualMeshCertificateSigningRequestTransitionFunction instructs the VirtualMeshCertificateSigningRequestWriter how to transition between an existing
// VirtualMeshCertificateSigningRequest object and a desired on an Upsert
type VirtualMeshCertificateSigningRequestTransitionFunction func(existing, desired *VirtualMeshCertificateSigningRequest) error

// Writer knows how to create, delete, and update VirtualMeshCertificateSigningRequests.
type VirtualMeshCertificateSigningRequestWriter interface {
	// Create saves the VirtualMeshCertificateSigningRequest object.
	CreateVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.CreateOption) error

	// Delete deletes the VirtualMeshCertificateSigningRequest object.
	DeleteVirtualMeshCertificateSigningRequest(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualMeshCertificateSigningRequest object.
	UpdateVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMeshCertificateSigningRequest object.
	PatchVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualMeshCertificateSigningRequest objects matching the given options.
	DeleteAllOfVirtualMeshCertificateSigningRequest(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualMeshCertificateSigningRequest object.
	UpsertVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, transitionFuncs ...VirtualMeshCertificateSigningRequestTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualMeshCertificateSigningRequest object.
type VirtualMeshCertificateSigningRequestStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualMeshCertificateSigningRequest object.
	UpdateVirtualMeshCertificateSigningRequestStatus(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMeshCertificateSigningRequest object's subresource.
	PatchVirtualMeshCertificateSigningRequestStatus(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualMeshCertificateSigningRequests.
type VirtualMeshCertificateSigningRequestClient interface {
	VirtualMeshCertificateSigningRequestReader
	VirtualMeshCertificateSigningRequestWriter
	VirtualMeshCertificateSigningRequestStatusWriter
}

type virtualMeshCertificateSigningRequestClient struct {
	client client.Client
}

func NewVirtualMeshCertificateSigningRequestClient(client client.Client) *virtualMeshCertificateSigningRequestClient {
	return &virtualMeshCertificateSigningRequestClient{client: client}
}

func (c *virtualMeshCertificateSigningRequestClient) GetVirtualMeshCertificateSigningRequest(ctx context.Context, key client.ObjectKey) (*VirtualMeshCertificateSigningRequest, error) {
	obj := &VirtualMeshCertificateSigningRequest{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualMeshCertificateSigningRequestClient) ListVirtualMeshCertificateSigningRequest(ctx context.Context, opts ...client.ListOption) (*VirtualMeshCertificateSigningRequestList, error) {
	list := &VirtualMeshCertificateSigningRequestList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualMeshCertificateSigningRequestClient) CreateVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) DeleteVirtualMeshCertificateSigningRequest(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualMeshCertificateSigningRequest{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) UpdateVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) PatchVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) DeleteAllOfVirtualMeshCertificateSigningRequest(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualMeshCertificateSigningRequest{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) UpsertVirtualMeshCertificateSigningRequest(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, transitionFuncs ...VirtualMeshCertificateSigningRequestTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*VirtualMeshCertificateSigningRequest), desired.(*VirtualMeshCertificateSigningRequest)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualMeshCertificateSigningRequestClient) UpdateVirtualMeshCertificateSigningRequestStatus(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualMeshCertificateSigningRequestClient) PatchVirtualMeshCertificateSigningRequestStatus(ctx context.Context, obj *VirtualMeshCertificateSigningRequest, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
