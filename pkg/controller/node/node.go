package node

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	ctldiskv1 "github.com/harvester/node-disk-manager/pkg/generated/controllers/harvesterhci.io/v1beta1"
	//ctllonghornv1 "github.com/harvester/node-disk-manager/pkg/generated/controllers/longhorn.io/v1beta1"
	ctlrookcephv1 "github.com/harvester/node-disk-manager/pkg/generated/controllers/ceph.rook.io/v1"
	"github.com/harvester/node-disk-manager/pkg/option"

	//longhornv1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	rookcephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
)

type Controller struct {
	namespace string

	BlockDevices     ctldiskv1.BlockDeviceController
	BlockDeviceCache ctldiskv1.BlockDeviceCache
	Nodes            ctlrookcephv1.CephClusterController
}

const (
	blockDeviceNodeHandlerName = "harvester-ndm-node-handler"
)

// Register register the longhorn node CRD controller
func Register(ctx context.Context, nodes ctlrookcephv1.CephClusterController, bds ctldiskv1.BlockDeviceController, opt *option.Option) error {

	c := &Controller{
		namespace:        opt.Namespace,
		Nodes:            nodes,
		BlockDevices:     bds,
		BlockDeviceCache: bds.Cache(),
	}

	nodes.OnRemove(ctx, blockDeviceNodeHandlerName, c.OnNodeDelete)
	return nil
}

// OnNodeDelete watch the node CR on remove and delete node related block devices
func (c *Controller) OnNodeDelete(key string, node *rookcephv1.CephCluster) (*rookcephv1.CephCluster, error) {
	if node == nil {
		return nil, nil
	}

	bds, err := c.BlockDeviceCache.List(c.namespace, labels.SelectorFromSet(map[string]string{
		v1.LabelHostname: node.Name,
	}))
	if err != nil {
		return node, err
	}

	for _, bd := range bds {
		if err := c.BlockDevices.Delete(c.namespace, bd.Name, &metav1.DeleteOptions{}); err != nil {
			return node, err
		}
	}
	return nil, nil
}
