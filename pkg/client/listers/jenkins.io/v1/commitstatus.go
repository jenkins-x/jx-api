// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CommitStatusLister helps list CommitStatuses.
type CommitStatusLister interface {
	// List lists all CommitStatuses in the indexer.
	List(selector labels.Selector) (ret []*v1.CommitStatus, err error)
	// CommitStatuses returns an object that can list and get CommitStatuses.
	CommitStatuses(namespace string) CommitStatusNamespaceLister
	CommitStatusListerExpansion
}

// commitStatusLister implements the CommitStatusLister interface.
type commitStatusLister struct {
	indexer cache.Indexer
}

// NewCommitStatusLister returns a new CommitStatusLister.
func NewCommitStatusLister(indexer cache.Indexer) CommitStatusLister {
	return &commitStatusLister{indexer: indexer}
}

// List lists all CommitStatuses in the indexer.
func (s *commitStatusLister) List(selector labels.Selector) (ret []*v1.CommitStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CommitStatus))
	})
	return ret, err
}

// CommitStatuses returns an object that can list and get CommitStatuses.
func (s *commitStatusLister) CommitStatuses(namespace string) CommitStatusNamespaceLister {
	return commitStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CommitStatusNamespaceLister helps list and get CommitStatuses.
type CommitStatusNamespaceLister interface {
	// List lists all CommitStatuses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CommitStatus, err error)
	// Get retrieves the CommitStatus from the indexer for a given namespace and name.
	Get(name string) (*v1.CommitStatus, error)
	CommitStatusNamespaceListerExpansion
}

// commitStatusNamespaceLister implements the CommitStatusNamespaceLister
// interface.
type commitStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CommitStatuses in the indexer for a given namespace.
func (s commitStatusNamespaceLister) List(selector labels.Selector) (ret []*v1.CommitStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CommitStatus))
	})
	return ret, err
}

// Get retrieves the CommitStatus from the indexer for a given namespace and name.
func (s commitStatusNamespaceLister) Get(name string) (*v1.CommitStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("commitstatus"), name)
	}
	return obj.(*v1.CommitStatus), nil
}
