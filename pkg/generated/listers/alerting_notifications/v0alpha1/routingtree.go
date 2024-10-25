// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by lister-gen. DO NOT EDIT.

package v0alpha1

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apis/alerting_notifications/v0alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// RoutingTreeLister helps list RoutingTrees.
// All objects returned here must be treated as read-only.
type RoutingTreeLister interface {
	// List lists all RoutingTrees in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.RoutingTree, err error)
	// RoutingTrees returns an object that can list and get RoutingTrees.
	RoutingTrees(namespace string) RoutingTreeNamespaceLister
	RoutingTreeListerExpansion
}

// routingTreeLister implements the RoutingTreeLister interface.
type routingTreeLister struct {
	listers.ResourceIndexer[*v0alpha1.RoutingTree]
}

// NewRoutingTreeLister returns a new RoutingTreeLister.
func NewRoutingTreeLister(indexer cache.Indexer) RoutingTreeLister {
	return &routingTreeLister{listers.New[*v0alpha1.RoutingTree](indexer, v0alpha1.Resource("routingtree"))}
}

// RoutingTrees returns an object that can list and get RoutingTrees.
func (s *routingTreeLister) RoutingTrees(namespace string) RoutingTreeNamespaceLister {
	return routingTreeNamespaceLister{listers.NewNamespaced[*v0alpha1.RoutingTree](s.ResourceIndexer, namespace)}
}

// RoutingTreeNamespaceLister helps list and get RoutingTrees.
// All objects returned here must be treated as read-only.
type RoutingTreeNamespaceLister interface {
	// List lists all RoutingTrees in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v0alpha1.RoutingTree, err error)
	// Get retrieves the RoutingTree from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v0alpha1.RoutingTree, error)
	RoutingTreeNamespaceListerExpansion
}

// routingTreeNamespaceLister implements the RoutingTreeNamespaceLister
// interface.
type routingTreeNamespaceLister struct {
	listers.ResourceIndexer[*v0alpha1.RoutingTree]
}
