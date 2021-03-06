/*
Copyright 2019 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/rancher/rio/pkg/apis/admin.rio.cattle.io/v1"
	scheme "github.com/rancher/rio/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FeaturesGetter has a method to return a FeatureInterface.
// A group's client should implement this interface.
type FeaturesGetter interface {
	Features(namespace string) FeatureInterface
}

// FeatureInterface has methods to work with Feature resources.
type FeatureInterface interface {
	Create(*v1.Feature) (*v1.Feature, error)
	Update(*v1.Feature) (*v1.Feature, error)
	UpdateStatus(*v1.Feature) (*v1.Feature, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Feature, error)
	List(opts metav1.ListOptions) (*v1.FeatureList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Feature, err error)
	FeatureExpansion
}

// features implements FeatureInterface
type features struct {
	client rest.Interface
	ns     string
}

// newFeatures returns a Features
func newFeatures(c *AdminV1Client, namespace string) *features {
	return &features{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the feature, and returns the corresponding feature object, and an error if there is any.
func (c *features) Get(name string, options metav1.GetOptions) (result *v1.Feature, err error) {
	result = &v1.Feature{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("features").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Features that match those selectors.
func (c *features) List(opts metav1.ListOptions) (result *v1.FeatureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.FeatureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("features").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested features.
func (c *features) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("features").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a feature and creates it.  Returns the server's representation of the feature, and an error, if there is any.
func (c *features) Create(feature *v1.Feature) (result *v1.Feature, err error) {
	result = &v1.Feature{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("features").
		Body(feature).
		Do().
		Into(result)
	return
}

// Update takes the representation of a feature and updates it. Returns the server's representation of the feature, and an error, if there is any.
func (c *features) Update(feature *v1.Feature) (result *v1.Feature, err error) {
	result = &v1.Feature{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("features").
		Name(feature.Name).
		Body(feature).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *features) UpdateStatus(feature *v1.Feature) (result *v1.Feature, err error) {
	result = &v1.Feature{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("features").
		Name(feature.Name).
		SubResource("status").
		Body(feature).
		Do().
		Into(result)
	return
}

// Delete takes name of the feature and deletes it. Returns an error if one occurs.
func (c *features) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("features").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *features) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("features").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched feature.
func (c *features) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Feature, err error) {
	result = &v1.Feature{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("features").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
