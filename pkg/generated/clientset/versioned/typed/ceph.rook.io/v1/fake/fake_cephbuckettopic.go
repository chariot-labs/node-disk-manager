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

// FakeCephBucketTopics implements CephBucketTopicInterface
type FakeCephBucketTopics struct {
	Fake *FakeCephV1
	ns   string
}

var cephbuckettopicsResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephbuckettopics"}

var cephbuckettopicsKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1", Kind: "CephBucketTopic"}

// Get takes name of the cephBucketTopic, and returns the corresponding cephBucketTopic object, and an error if there is any.
func (c *FakeCephBucketTopics) Get(ctx context.Context, name string, options v1.GetOptions) (result *cephrookiov1.CephBucketTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephbuckettopicsResource, c.ns, name), &cephrookiov1.CephBucketTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketTopic), err
}

// List takes label and field selectors, and returns the list of CephBucketTopics that match those selectors.
func (c *FakeCephBucketTopics) List(ctx context.Context, opts v1.ListOptions) (result *cephrookiov1.CephBucketTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephbuckettopicsResource, cephbuckettopicsKind, c.ns, opts), &cephrookiov1.CephBucketTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cephrookiov1.CephBucketTopicList{ListMeta: obj.(*cephrookiov1.CephBucketTopicList).ListMeta}
	for _, item := range obj.(*cephrookiov1.CephBucketTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephBucketTopics.
func (c *FakeCephBucketTopics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephbuckettopicsResource, c.ns, opts))

}

// Create takes the representation of a cephBucketTopic and creates it.  Returns the server's representation of the cephBucketTopic, and an error, if there is any.
func (c *FakeCephBucketTopics) Create(ctx context.Context, cephBucketTopic *cephrookiov1.CephBucketTopic, opts v1.CreateOptions) (result *cephrookiov1.CephBucketTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephbuckettopicsResource, c.ns, cephBucketTopic), &cephrookiov1.CephBucketTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketTopic), err
}

// Update takes the representation of a cephBucketTopic and updates it. Returns the server's representation of the cephBucketTopic, and an error, if there is any.
func (c *FakeCephBucketTopics) Update(ctx context.Context, cephBucketTopic *cephrookiov1.CephBucketTopic, opts v1.UpdateOptions) (result *cephrookiov1.CephBucketTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephbuckettopicsResource, c.ns, cephBucketTopic), &cephrookiov1.CephBucketTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketTopic), err
}

// Delete takes name of the cephBucketTopic and deletes it. Returns an error if one occurs.
func (c *FakeCephBucketTopics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cephbuckettopicsResource, c.ns, name), &cephrookiov1.CephBucketTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephBucketTopics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephbuckettopicsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cephrookiov1.CephBucketTopicList{})
	return err
}

// Patch applies the patch and returns the patched cephBucketTopic.
func (c *FakeCephBucketTopics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cephrookiov1.CephBucketTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephbuckettopicsResource, c.ns, name, pt, data, subresources...), &cephrookiov1.CephBucketTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephBucketTopic), err
}
