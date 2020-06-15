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

// SchedulersGetter has a method to return a SchedulerInterface.
// A group's client should implement this interface.
type SchedulersGetter interface {
	Schedulers(namespace string) SchedulerInterface
}

// SchedulerInterface has methods to work with Scheduler resources.
type SchedulerInterface interface {
	Create(*v1.Scheduler) (*v1.Scheduler, error)
	Update(*v1.Scheduler) (*v1.Scheduler, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Scheduler, error)
	List(opts metav1.ListOptions) (*v1.SchedulerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scheduler, err error)
	SchedulerExpansion
}

// schedulers implements SchedulerInterface
type schedulers struct {
	client rest.Interface
	ns     string
}

// newSchedulers returns a Schedulers
func newSchedulers(c *JenkinsV1Client, namespace string) *schedulers {
	return &schedulers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scheduler, and returns the corresponding scheduler object, and an error if there is any.
func (c *schedulers) Get(name string, options metav1.GetOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Schedulers that match those selectors.
func (c *schedulers) List(opts metav1.ListOptions) (result *v1.SchedulerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SchedulerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulers.
func (c *schedulers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a scheduler and creates it.  Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Create(scheduler *v1.Scheduler) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schedulers").
		Body(scheduler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scheduler and updates it. Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Update(scheduler *v1.Scheduler) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulers").
		Name(scheduler.Name).
		Body(scheduler).
		Do().
		Into(result)
	return
}

// Delete takes name of the scheduler and deletes it. Returns an error if one occurs.
func (c *schedulers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scheduler.
func (c *schedulers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schedulers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
