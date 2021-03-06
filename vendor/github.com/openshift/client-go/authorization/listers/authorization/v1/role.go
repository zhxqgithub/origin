// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoleLister helps list Roles.
type RoleLister interface {
	// List lists all Roles in the indexer.
	List(selector labels.Selector) (ret []*v1.Role, err error)
	// Roles returns an object that can list and get Roles.
	Roles(namespace string) RoleNamespaceLister
	RoleListerExpansion
}

// roleLister implements the RoleLister interface.
type roleLister struct {
	indexer cache.Indexer
}

// NewRoleLister returns a new RoleLister.
func NewRoleLister(indexer cache.Indexer) RoleLister {
	return &roleLister{indexer: indexer}
}

// List lists all Roles in the indexer.
func (s *roleLister) List(selector labels.Selector) (ret []*v1.Role, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Role))
	})
	return ret, err
}

// Roles returns an object that can list and get Roles.
func (s *roleLister) Roles(namespace string) RoleNamespaceLister {
	return roleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoleNamespaceLister helps list and get Roles.
type RoleNamespaceLister interface {
	// List lists all Roles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Role, err error)
	// Get retrieves the Role from the indexer for a given namespace and name.
	Get(name string) (*v1.Role, error)
	RoleNamespaceListerExpansion
}

// roleNamespaceLister implements the RoleNamespaceLister
// interface.
type roleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Roles in the indexer for a given namespace.
func (s roleNamespaceLister) List(selector labels.Selector) (ret []*v1.Role, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Role))
	})
	return ret, err
}

// Get retrieves the Role from the indexer for a given namespace and name.
func (s roleNamespaceLister) Get(name string) (*v1.Role, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("role"), name)
	}
	return obj.(*v1.Role), nil
}
