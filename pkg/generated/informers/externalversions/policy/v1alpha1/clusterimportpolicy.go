// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	policyv1alpha1 "github.com/clusterpedia-io/api/policy/v1alpha1"
	versioned "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/clusterpedia-io/clusterpedia/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/clusterpedia-io/clusterpedia/pkg/generated/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterImportPolicyInformer provides access to a shared informer and lister for
// ClusterImportPolicies.
type ClusterImportPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterImportPolicyLister
}

type clusterImportPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterImportPolicyInformer constructs a new informer for ClusterImportPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterImportPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterImportPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterImportPolicyInformer constructs a new informer for ClusterImportPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterImportPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterImportPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().ClusterImportPolicies().Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.ClusterImportPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterImportPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterImportPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterImportPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.ClusterImportPolicy{}, f.defaultInformer)
}

func (f *clusterImportPolicyInformer) Lister() v1alpha1.ClusterImportPolicyLister {
	return v1alpha1.NewClusterImportPolicyLister(f.Informer().GetIndexer())
}