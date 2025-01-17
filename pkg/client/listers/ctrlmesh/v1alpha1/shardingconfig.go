/*
Copyright 2023 The KusionStack Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/KusionStack/ctrlmesh/pkg/apis/ctrlmesh/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ShardingConfigLister helps list ShardingConfigs.
// All objects returned here must be treated as read-only.
type ShardingConfigLister interface {
	// List lists all ShardingConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShardingConfig, err error)
	// ShardingConfigs returns an object that can list and get ShardingConfigs.
	ShardingConfigs(namespace string) ShardingConfigNamespaceLister
	ShardingConfigListerExpansion
}

// shardingConfigLister implements the ShardingConfigLister interface.
type shardingConfigLister struct {
	indexer cache.Indexer
}

// NewShardingConfigLister returns a new ShardingConfigLister.
func NewShardingConfigLister(indexer cache.Indexer) ShardingConfigLister {
	return &shardingConfigLister{indexer: indexer}
}

// List lists all ShardingConfigs in the indexer.
func (s *shardingConfigLister) List(selector labels.Selector) (ret []*v1alpha1.ShardingConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShardingConfig))
	})
	return ret, err
}

// ShardingConfigs returns an object that can list and get ShardingConfigs.
func (s *shardingConfigLister) ShardingConfigs(namespace string) ShardingConfigNamespaceLister {
	return shardingConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShardingConfigNamespaceLister helps list and get ShardingConfigs.
// All objects returned here must be treated as read-only.
type ShardingConfigNamespaceLister interface {
	// List lists all ShardingConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShardingConfig, err error)
	// Get retrieves the ShardingConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShardingConfig, error)
	ShardingConfigNamespaceListerExpansion
}

// shardingConfigNamespaceLister implements the ShardingConfigNamespaceLister
// interface.
type shardingConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShardingConfigs in the indexer for a given namespace.
func (s shardingConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShardingConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShardingConfig))
	})
	return ret, err
}

// Get retrieves the ShardingConfig from the indexer for a given namespace and name.
func (s shardingConfigNamespaceLister) Get(name string) (*v1alpha1.ShardingConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shardingconfig"), name)
	}
	return obj.(*v1alpha1.ShardingConfig), nil
}
