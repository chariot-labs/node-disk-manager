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

// FakeCephFilesystemSubVolumeGroups implements CephFilesystemSubVolumeGroupInterface
type FakeCephFilesystemSubVolumeGroups struct {
	Fake *FakeCephV1
	ns   string
}

var cephfilesystemsubvolumegroupsResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephfilesystemsubvolumegroups"}

var cephfilesystemsubvolumegroupsKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1", Kind: "CephFilesystemSubVolumeGroup"}

// Get takes name of the cephFilesystemSubVolumeGroup, and returns the corresponding cephFilesystemSubVolumeGroup object, and an error if there is any.
func (c *FakeCephFilesystemSubVolumeGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *cephrookiov1.CephFilesystemSubVolumeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephfilesystemsubvolumegroupsResource, c.ns, name), &cephrookiov1.CephFilesystemSubVolumeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystemSubVolumeGroup), err
}

// List takes label and field selectors, and returns the list of CephFilesystemSubVolumeGroups that match those selectors.
func (c *FakeCephFilesystemSubVolumeGroups) List(ctx context.Context, opts v1.ListOptions) (result *cephrookiov1.CephFilesystemSubVolumeGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephfilesystemsubvolumegroupsResource, cephfilesystemsubvolumegroupsKind, c.ns, opts), &cephrookiov1.CephFilesystemSubVolumeGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cephrookiov1.CephFilesystemSubVolumeGroupList{ListMeta: obj.(*cephrookiov1.CephFilesystemSubVolumeGroupList).ListMeta}
	for _, item := range obj.(*cephrookiov1.CephFilesystemSubVolumeGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephFilesystemSubVolumeGroups.
func (c *FakeCephFilesystemSubVolumeGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephfilesystemsubvolumegroupsResource, c.ns, opts))

}

// Create takes the representation of a cephFilesystemSubVolumeGroup and creates it.  Returns the server's representation of the cephFilesystemSubVolumeGroup, and an error, if there is any.
func (c *FakeCephFilesystemSubVolumeGroups) Create(ctx context.Context, cephFilesystemSubVolumeGroup *cephrookiov1.CephFilesystemSubVolumeGroup, opts v1.CreateOptions) (result *cephrookiov1.CephFilesystemSubVolumeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephfilesystemsubvolumegroupsResource, c.ns, cephFilesystemSubVolumeGroup), &cephrookiov1.CephFilesystemSubVolumeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystemSubVolumeGroup), err
}

// Update takes the representation of a cephFilesystemSubVolumeGroup and updates it. Returns the server's representation of the cephFilesystemSubVolumeGroup, and an error, if there is any.
func (c *FakeCephFilesystemSubVolumeGroups) Update(ctx context.Context, cephFilesystemSubVolumeGroup *cephrookiov1.CephFilesystemSubVolumeGroup, opts v1.UpdateOptions) (result *cephrookiov1.CephFilesystemSubVolumeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephfilesystemsubvolumegroupsResource, c.ns, cephFilesystemSubVolumeGroup), &cephrookiov1.CephFilesystemSubVolumeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystemSubVolumeGroup), err
}

// Delete takes name of the cephFilesystemSubVolumeGroup and deletes it. Returns an error if one occurs.
func (c *FakeCephFilesystemSubVolumeGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cephfilesystemsubvolumegroupsResource, c.ns, name), &cephrookiov1.CephFilesystemSubVolumeGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephFilesystemSubVolumeGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephfilesystemsubvolumegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cephrookiov1.CephFilesystemSubVolumeGroupList{})
	return err
}

// Patch applies the patch and returns the patched cephFilesystemSubVolumeGroup.
func (c *FakeCephFilesystemSubVolumeGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cephrookiov1.CephFilesystemSubVolumeGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephfilesystemsubvolumegroupsResource, c.ns, name, pt, data, subresources...), &cephrookiov1.CephFilesystemSubVolumeGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystemSubVolumeGroup), err
}
