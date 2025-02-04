/*
Copyright 2022 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	cephrookiov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephBucketNotifications implements CephBucketNotificationInterface
type FakeCephBucketNotifications struct {
	Fake *FakeCephV1
	ns   string
}

var cephbucketnotificationsResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephbucketnotifications"}

var cephbucketnotificationsKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1", Kind: "CephBucketNotification"}

// Get takes name of the cephBucketNotification, and returns the corresponding cephBucketNotification object, and an error if there is any.
func (c *FakeCephBucketNotifications) Get(ctx context.Context, name string, options v1.GetOptions) (result *cephrookiov1.CephBucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephbucketnotificationsResource, c.ns, name), &cephrookiov1.CephBucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketNotification), err
}

// List takes label and field selectors, and returns the list of CephBucketNotifications that match those selectors.
func (c *FakeCephBucketNotifications) List(ctx context.Context, opts v1.ListOptions) (result *cephrookiov1.CephBucketNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephbucketnotificationsResource, cephbucketnotificationsKind, c.ns, opts), &cephrookiov1.CephBucketNotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cephrookiov1.CephBucketNotificationList{ListMeta: obj.(*cephrookiov1.CephBucketNotificationList).ListMeta}
	for _, item := range obj.(*cephrookiov1.CephBucketNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephBucketNotifications.
func (c *FakeCephBucketNotifications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephbucketnotificationsResource, c.ns, opts))

}

// Create takes the representation of a cephBucketNotification and creates it.  Returns the server's representation of the cephBucketNotification, and an error, if there is any.
func (c *FakeCephBucketNotifications) Create(ctx context.Context, cephBucketNotification *cephrookiov1.CephBucketNotification, opts v1.CreateOptions) (result *cephrookiov1.CephBucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephbucketnotificationsResource, c.ns, cephBucketNotification), &cephrookiov1.CephBucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketNotification), err
}

// Update takes the representation of a cephBucketNotification and updates it. Returns the server's representation of the cephBucketNotification, and an error, if there is any.
func (c *FakeCephBucketNotifications) Update(ctx context.Context, cephBucketNotification *cephrookiov1.CephBucketNotification, opts v1.UpdateOptions) (result *cephrookiov1.CephBucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephbucketnotificationsResource, c.ns, cephBucketNotification), &cephrookiov1.CephBucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketNotification), err
}

// Delete takes name of the cephBucketNotification and deletes it. Returns an error if one occurs.
func (c *FakeCephBucketNotifications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cephbucketnotificationsResource, c.ns, name), &cephrookiov1.CephBucketNotification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephBucketNotifications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephbucketnotificationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cephrookiov1.CephBucketNotificationList{})
	return err
}

// Patch applies the patch and returns the patched cephBucketNotification.
func (c *FakeCephBucketNotifications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cephrookiov1.CephBucketNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephbucketnotificationsResource, c.ns, name, pt, data, subresources...), &cephrookiov1.CephBucketNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketNotification), err
}
