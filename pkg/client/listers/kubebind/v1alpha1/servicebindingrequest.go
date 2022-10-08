/*
Copyright 2022 The Kubectl Bind contributors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// ServiceBindingRequestLister helps list ServiceBindingRequests.
// All objects returned here must be treated as read-only.
type ServiceBindingRequestLister interface {
	// List lists all ServiceBindingRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingRequest, err error)
	// ServiceBindingRequests returns an object that can list and get ServiceBindingRequests.
	ServiceBindingRequests(namespace string) ServiceBindingRequestNamespaceLister
	ServiceBindingRequestListerExpansion
}

// serviceBindingRequestLister implements the ServiceBindingRequestLister interface.
type serviceBindingRequestLister struct {
	indexer cache.Indexer
}

// NewServiceBindingRequestLister returns a new ServiceBindingRequestLister.
func NewServiceBindingRequestLister(indexer cache.Indexer) ServiceBindingRequestLister {
	return &serviceBindingRequestLister{indexer: indexer}
}

// List lists all ServiceBindingRequests in the indexer.
func (s *serviceBindingRequestLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceBindingRequest))
	})
	return ret, err
}

// ServiceBindingRequests returns an object that can list and get ServiceBindingRequests.
func (s *serviceBindingRequestLister) ServiceBindingRequests(namespace string) ServiceBindingRequestNamespaceLister {
	return serviceBindingRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceBindingRequestNamespaceLister helps list and get ServiceBindingRequests.
// All objects returned here must be treated as read-only.
type ServiceBindingRequestNamespaceLister interface {
	// List lists all ServiceBindingRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingRequest, err error)
	// Get retrieves the ServiceBindingRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceBindingRequest, error)
	ServiceBindingRequestNamespaceListerExpansion
}

// serviceBindingRequestNamespaceLister implements the ServiceBindingRequestNamespaceLister
// interface.
type serviceBindingRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceBindingRequests in the indexer for a given namespace.
func (s serviceBindingRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceBindingRequest))
	})
	return ret, err
}

// Get retrieves the ServiceBindingRequest from the indexer for a given namespace and name.
func (s serviceBindingRequestNamespaceLister) Get(name string) (*v1alpha1.ServiceBindingRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicebindingrequest"), name)
	}
	return obj.(*v1alpha1.ServiceBindingRequest), nil
}
