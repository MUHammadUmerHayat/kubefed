/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/marun/federation-v2/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedJobLister helps list FederatedJobs.
type FederatedJobLister interface {
	// List lists all FederatedJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedJob, err error)
	// FederatedJobs returns an object that can list and get FederatedJobs.
	FederatedJobs(namespace string) FederatedJobNamespaceLister
	FederatedJobListerExpansion
}

// federatedJobLister implements the FederatedJobLister interface.
type federatedJobLister struct {
	indexer cache.Indexer
}

// NewFederatedJobLister returns a new FederatedJobLister.
func NewFederatedJobLister(indexer cache.Indexer) FederatedJobLister {
	return &federatedJobLister{indexer: indexer}
}

// List lists all FederatedJobs in the indexer.
func (s *federatedJobLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedJob))
	})
	return ret, err
}

// FederatedJobs returns an object that can list and get FederatedJobs.
func (s *federatedJobLister) FederatedJobs(namespace string) FederatedJobNamespaceLister {
	return federatedJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedJobNamespaceLister helps list and get FederatedJobs.
type FederatedJobNamespaceLister interface {
	// List lists all FederatedJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedJob, err error)
	// Get retrieves the FederatedJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedJob, error)
	FederatedJobNamespaceListerExpansion
}

// federatedJobNamespaceLister implements the FederatedJobNamespaceLister
// interface.
type federatedJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedJobs in the indexer for a given namespace.
func (s federatedJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedJob))
	})
	return ret, err
}

// Get retrieves the FederatedJob from the indexer for a given namespace and name.
func (s federatedJobNamespaceLister) Get(name string) (*v1alpha1.FederatedJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedjob"), name)
	}
	return obj.(*v1alpha1.FederatedJob), nil
}
