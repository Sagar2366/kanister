// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=cr, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("actionsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cr().V1alpha1().ActionSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("blueprints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cr().V1alpha1().Blueprints().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("profiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cr().V1alpha1().Profiles().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
