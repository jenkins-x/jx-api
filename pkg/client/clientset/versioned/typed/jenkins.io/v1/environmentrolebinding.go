// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx-api/v1/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnvironmentRoleBindingsGetter has a method to return a EnvironmentRoleBindingInterface.
// A group's client should implement this interface.
type EnvironmentRoleBindingsGetter interface {
	EnvironmentRoleBindings(namespace string) EnvironmentRoleBindingInterface
}

// EnvironmentRoleBindingInterface has methods to work with EnvironmentRoleBinding resources.
type EnvironmentRoleBindingInterface interface {
	Create(*v1.EnvironmentRoleBinding) (*v1.EnvironmentRoleBinding, error)
	Update(*v1.EnvironmentRoleBinding) (*v1.EnvironmentRoleBinding, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.EnvironmentRoleBinding, error)
	List(opts metav1.ListOptions) (*v1.EnvironmentRoleBindingList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EnvironmentRoleBinding, err error)
	EnvironmentRoleBindingExpansion
}

// environmentRoleBindings implements EnvironmentRoleBindingInterface
type environmentRoleBindings struct {
	client rest.Interface
	ns     string
}

// newEnvironmentRoleBindings returns a EnvironmentRoleBindings
func newEnvironmentRoleBindings(c *JenkinsV1Client, namespace string) *environmentRoleBindings {
	return &environmentRoleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the environmentRoleBinding, and returns the corresponding environmentRoleBinding object, and an error if there is any.
func (c *environmentRoleBindings) Get(name string, options metav1.GetOptions) (result *v1.EnvironmentRoleBinding, err error) {
	result = &v1.EnvironmentRoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EnvironmentRoleBindings that match those selectors.
func (c *environmentRoleBindings) List(opts metav1.ListOptions) (result *v1.EnvironmentRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EnvironmentRoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested environmentRoleBindings.
func (c *environmentRoleBindings) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a environmentRoleBinding and creates it.  Returns the server's representation of the environmentRoleBinding, and an error, if there is any.
func (c *environmentRoleBindings) Create(environmentRoleBinding *v1.EnvironmentRoleBinding) (result *v1.EnvironmentRoleBinding, err error) {
	result = &v1.EnvironmentRoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		Body(environmentRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a environmentRoleBinding and updates it. Returns the server's representation of the environmentRoleBinding, and an error, if there is any.
func (c *environmentRoleBindings) Update(environmentRoleBinding *v1.EnvironmentRoleBinding) (result *v1.EnvironmentRoleBinding, err error) {
	result = &v1.EnvironmentRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		Name(environmentRoleBinding.Name).
		Body(environmentRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the environmentRoleBinding and deletes it. Returns an error if one occurs.
func (c *environmentRoleBindings) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *environmentRoleBindings) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environmentrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched environmentRoleBinding.
func (c *environmentRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EnvironmentRoleBinding, err error) {
	result = &v1.EnvironmentRoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("environmentrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
