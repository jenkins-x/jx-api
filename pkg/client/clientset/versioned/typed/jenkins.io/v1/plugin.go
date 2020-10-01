/*
Copyright 2020 The Jenkins X Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PluginsGetter has a method to return a PluginInterface.
// A group's client should implement this interface.
type PluginsGetter interface {
	Plugins(namespace string) PluginInterface
}

// PluginInterface has methods to work with Plugin resources.
type PluginInterface interface {
	Create(ctx context.Context, plugin *v1.Plugin, opts metav1.CreateOptions) (*v1.Plugin, error)
	Update(ctx context.Context, plugin *v1.Plugin, opts metav1.UpdateOptions) (*v1.Plugin, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Plugin, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PluginList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Plugin, err error)
	PluginExpansion
}

// plugins implements PluginInterface
type plugins struct {
	client rest.Interface
	ns     string
}

// newPlugins returns a Plugins
func newPlugins(c *JenkinsV1Client, namespace string) *plugins {
	return &plugins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the plugin, and returns the corresponding plugin object, and an error if there is any.
func (c *plugins) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Plugins that match those selectors.
func (c *plugins) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PluginList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PluginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested plugins.
func (c *plugins) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a plugin and creates it.  Returns the server's representation of the plugin, and an error, if there is any.
func (c *plugins) Create(ctx context.Context, plugin *v1.Plugin, opts metav1.CreateOptions) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(plugin).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a plugin and updates it. Returns the server's representation of the plugin, and an error, if there is any.
func (c *plugins) Update(ctx context.Context, plugin *v1.Plugin, opts metav1.UpdateOptions) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("plugins").
		Name(plugin.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(plugin).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the plugin and deletes it. Returns an error if one occurs.
func (c *plugins) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("plugins").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *plugins) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched plugin.
func (c *plugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("plugins").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
