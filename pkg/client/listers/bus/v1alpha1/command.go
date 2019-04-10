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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/kube-batch/pkg/apis/bus/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CommandLister helps list Commands.
type CommandLister interface {
	// List lists all Commands in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Command, err error)
	// Commands returns an object that can list and get Commands.
	Commands(namespace string) CommandNamespaceLister
	CommandListerExpansion
}

// commandLister implements the CommandLister interface.
type commandLister struct {
	indexer cache.Indexer
}

// NewCommandLister returns a new CommandLister.
func NewCommandLister(indexer cache.Indexer) CommandLister {
	return &commandLister{indexer: indexer}
}

// List lists all Commands in the indexer.
func (s *commandLister) List(selector labels.Selector) (ret []*v1alpha1.Command, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Command))
	})
	return ret, err
}

// Commands returns an object that can list and get Commands.
func (s *commandLister) Commands(namespace string) CommandNamespaceLister {
	return commandNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CommandNamespaceLister helps list and get Commands.
type CommandNamespaceLister interface {
	// List lists all Commands in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Command, err error)
	// Get retrieves the Command from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Command, error)
	CommandNamespaceListerExpansion
}

// commandNamespaceLister implements the CommandNamespaceLister
// interface.
type commandNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Commands in the indexer for a given namespace.
func (s commandNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Command, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Command))
	})
	return ret, err
}

// Get retrieves the Command from the indexer for a given namespace and name.
func (s commandNamespaceLister) Get(name string) (*v1alpha1.Command, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("command"), name)
	}
	return obj.(*v1alpha1.Command), nil
}
