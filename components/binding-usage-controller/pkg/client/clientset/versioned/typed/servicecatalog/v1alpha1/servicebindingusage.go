// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	scheme "github.com/kyma-project/kyma/components/binding-usage-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceBindingUsagesGetter has a method to return a ServiceBindingUsageInterface.
// A group's client should implement this interface.
type ServiceBindingUsagesGetter interface {
	ServiceBindingUsages(namespace string) ServiceBindingUsageInterface
}

// ServiceBindingUsageInterface has methods to work with ServiceBindingUsage resources.
type ServiceBindingUsageInterface interface {
	Create(*v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error)
	Update(*v1alpha1.ServiceBindingUsage) (*v1alpha1.ServiceBindingUsage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceBindingUsage, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceBindingUsageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceBindingUsage, err error)
	ServiceBindingUsageExpansion
}

// serviceBindingUsages implements ServiceBindingUsageInterface
type serviceBindingUsages struct {
	client rest.Interface
	ns     string
}

// newServiceBindingUsages returns a ServiceBindingUsages
func newServiceBindingUsages(c *ServicecatalogV1alpha1Client, namespace string) *serviceBindingUsages {
	return &serviceBindingUsages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceBindingUsage, and returns the corresponding serviceBindingUsage object, and an error if there is any.
func (c *serviceBindingUsages) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceBindingUsage, err error) {
	result = &v1alpha1.ServiceBindingUsage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebindingusages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceBindingUsages that match those selectors.
func (c *serviceBindingUsages) List(opts v1.ListOptions) (result *v1alpha1.ServiceBindingUsageList, err error) {
	result = &v1alpha1.ServiceBindingUsageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebindingusages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceBindingUsages.
func (c *serviceBindingUsages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicebindingusages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a serviceBindingUsage and creates it.  Returns the server's representation of the serviceBindingUsage, and an error, if there is any.
func (c *serviceBindingUsages) Create(serviceBindingUsage *v1alpha1.ServiceBindingUsage) (result *v1alpha1.ServiceBindingUsage, err error) {
	result = &v1alpha1.ServiceBindingUsage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicebindingusages").
		Body(serviceBindingUsage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceBindingUsage and updates it. Returns the server's representation of the serviceBindingUsage, and an error, if there is any.
func (c *serviceBindingUsages) Update(serviceBindingUsage *v1alpha1.ServiceBindingUsage) (result *v1alpha1.ServiceBindingUsage, err error) {
	result = &v1alpha1.ServiceBindingUsage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebindingusages").
		Name(serviceBindingUsage.Name).
		Body(serviceBindingUsage).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceBindingUsage and deletes it. Returns an error if one occurs.
func (c *serviceBindingUsages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebindingusages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceBindingUsages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebindingusages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceBindingUsage.
func (c *serviceBindingUsages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceBindingUsage, err error) {
	result = &v1alpha1.ServiceBindingUsage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicebindingusages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}