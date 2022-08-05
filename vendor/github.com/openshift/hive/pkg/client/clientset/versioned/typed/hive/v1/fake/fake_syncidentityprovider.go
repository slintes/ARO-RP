// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	hivev1 "github.com/openshift/hive/apis/hive/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSyncIdentityProviders implements SyncIdentityProviderInterface
type FakeSyncIdentityProviders struct {
	Fake *FakeHiveV1
	ns   string
}

var syncidentityprovidersResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "syncidentityproviders"}

var syncidentityprovidersKind = schema.GroupVersionKind{Group: "hive.openshift.io", Version: "v1", Kind: "SyncIdentityProvider"}

// Get takes name of the syncIdentityProvider, and returns the corresponding syncIdentityProvider object, and an error if there is any.
func (c *FakeSyncIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *hivev1.SyncIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(syncidentityprovidersResource, c.ns, name), &hivev1.SyncIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncIdentityProvider), err
}

// List takes label and field selectors, and returns the list of SyncIdentityProviders that match those selectors.
func (c *FakeSyncIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *hivev1.SyncIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(syncidentityprovidersResource, syncidentityprovidersKind, c.ns, opts), &hivev1.SyncIdentityProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hivev1.SyncIdentityProviderList{ListMeta: obj.(*hivev1.SyncIdentityProviderList).ListMeta}
	for _, item := range obj.(*hivev1.SyncIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested syncIdentityProviders.
func (c *FakeSyncIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(syncidentityprovidersResource, c.ns, opts))

}

// Create takes the representation of a syncIdentityProvider and creates it.  Returns the server's representation of the syncIdentityProvider, and an error, if there is any.
func (c *FakeSyncIdentityProviders) Create(ctx context.Context, syncIdentityProvider *hivev1.SyncIdentityProvider, opts v1.CreateOptions) (result *hivev1.SyncIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(syncidentityprovidersResource, c.ns, syncIdentityProvider), &hivev1.SyncIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncIdentityProvider), err
}

// Update takes the representation of a syncIdentityProvider and updates it. Returns the server's representation of the syncIdentityProvider, and an error, if there is any.
func (c *FakeSyncIdentityProviders) Update(ctx context.Context, syncIdentityProvider *hivev1.SyncIdentityProvider, opts v1.UpdateOptions) (result *hivev1.SyncIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(syncidentityprovidersResource, c.ns, syncIdentityProvider), &hivev1.SyncIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSyncIdentityProviders) UpdateStatus(ctx context.Context, syncIdentityProvider *hivev1.SyncIdentityProvider, opts v1.UpdateOptions) (*hivev1.SyncIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(syncidentityprovidersResource, "status", c.ns, syncIdentityProvider), &hivev1.SyncIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncIdentityProvider), err
}

// Delete takes name of the syncIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeSyncIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(syncidentityprovidersResource, c.ns, name, opts), &hivev1.SyncIdentityProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSyncIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(syncidentityprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &hivev1.SyncIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched syncIdentityProvider.
func (c *FakeSyncIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hivev1.SyncIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(syncidentityprovidersResource, c.ns, name, pt, data, subresources...), &hivev1.SyncIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncIdentityProvider), err
}