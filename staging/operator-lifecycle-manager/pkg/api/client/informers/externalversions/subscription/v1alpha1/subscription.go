/*
Copyright 2018 CoreOS, Inc.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	subscription_v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/subscription/v1alpha1"
	versioned "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	internalinterfaces "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/listers/subscription/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// SubscriptionInformer provides access to a shared informer and lister for
// Subscriptions.
type SubscriptionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SubscriptionLister
}

type subscriptionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSubscriptionInformer constructs a new informer for Subscription type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSubscriptionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSubscriptionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSubscriptionInformer constructs a new informer for Subscription type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSubscriptionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SubscriptionV1alpha1().Subscriptions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SubscriptionV1alpha1().Subscriptions(namespace).Watch(options)
			},
		},
		&subscription_v1alpha1.Subscription{},
		resyncPeriod,
		indexers,
	)
}

func (f *subscriptionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSubscriptionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *subscriptionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&subscription_v1alpha1.Subscription{}, f.defaultInformer)
}

func (f *subscriptionInformer) Lister() v1alpha1.SubscriptionLister {
	return v1alpha1.NewSubscriptionLister(f.Informer().GetIndexer())
}
