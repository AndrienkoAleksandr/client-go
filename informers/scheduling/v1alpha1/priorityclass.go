/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	internalinterfaces "github.com/AndrienkoAleksandr/client-go/informers/internalinterfaces"
	kubernetes "github.com/AndrienkoAleksandr/client-go/kubernetes"
	v1alpha1 "github.com/AndrienkoAleksandr/client-go/listers/scheduling/v1alpha1"
	cache "github.com/AndrienkoAleksandr/client-go/tools/cache"
	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
)

// PriorityClassInformer provides access to a shared informer and lister for
// PriorityClasses.
type PriorityClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PriorityClassLister
}

type priorityClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPriorityClassInformer constructs a new informer for PriorityClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPriorityClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPriorityClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPriorityClassInformer constructs a new informer for PriorityClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPriorityClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1alpha1().PriorityClasses().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1alpha1().PriorityClasses().Watch(options)
			},
		},
		&schedulingv1alpha1.PriorityClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *priorityClassInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPriorityClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *priorityClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schedulingv1alpha1.PriorityClass{}, f.defaultInformer)
}

func (f *priorityClassInformer) Lister() v1alpha1.PriorityClassLister {
	return v1alpha1.NewPriorityClassLister(f.Informer().GetIndexer())
}
