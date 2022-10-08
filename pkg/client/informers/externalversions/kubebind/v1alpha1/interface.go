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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/kube-bind/kube-bind/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterBindings returns a ClusterBindingInformer.
	ClusterBindings() ClusterBindingInformer
	// ServiceBindings returns a ServiceBindingInformer.
	ServiceBindings() ServiceBindingInformer
	// ServiceBindingRequests returns a ServiceBindingRequestInformer.
	ServiceBindingRequests() ServiceBindingRequestInformer
	// ServiceExports returns a ServiceExportInformer.
	ServiceExports() ServiceExportInformer
	// ServiceExportResources returns a ServiceExportResourceInformer.
	ServiceExportResources() ServiceExportResourceInformer
	// ServiceNamespaces returns a ServiceNamespaceInformer.
	ServiceNamespaces() ServiceNamespaceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterBindings returns a ClusterBindingInformer.
func (v *version) ClusterBindings() ClusterBindingInformer {
	return &clusterBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceBindings returns a ServiceBindingInformer.
func (v *version) ServiceBindings() ServiceBindingInformer {
	return &serviceBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceBindingRequests returns a ServiceBindingRequestInformer.
func (v *version) ServiceBindingRequests() ServiceBindingRequestInformer {
	return &serviceBindingRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceExports returns a ServiceExportInformer.
func (v *version) ServiceExports() ServiceExportInformer {
	return &serviceExportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceExportResources returns a ServiceExportResourceInformer.
func (v *version) ServiceExportResources() ServiceExportResourceInformer {
	return &serviceExportResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceNamespaces returns a ServiceNamespaceInformer.
func (v *version) ServiceNamespaces() ServiceNamespaceInformer {
	return &serviceNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
