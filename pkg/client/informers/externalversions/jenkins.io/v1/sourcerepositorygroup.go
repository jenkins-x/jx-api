// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx-api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx-api/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SourceRepositoryGroupInformer provides access to a shared informer and lister for
// SourceRepositoryGroups.
type SourceRepositoryGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SourceRepositoryGroupLister
}

type sourceRepositoryGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSourceRepositoryGroupInformer constructs a new informer for SourceRepositoryGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSourceRepositoryGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSourceRepositoryGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSourceRepositoryGroupInformer constructs a new informer for SourceRepositoryGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSourceRepositoryGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().SourceRepositoryGroups(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().SourceRepositoryGroups(namespace).Watch(options)
			},
		},
		&jenkinsiov1.SourceRepositoryGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *sourceRepositoryGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSourceRepositoryGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sourceRepositoryGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.SourceRepositoryGroup{}, f.defaultInformer)
}

func (f *sourceRepositoryGroupInformer) Lister() v1.SourceRepositoryGroupLister {
	return v1.NewSourceRepositoryGroupLister(f.Informer().GetIndexer())
}
