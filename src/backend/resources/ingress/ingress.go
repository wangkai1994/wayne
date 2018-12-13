package ingress

import (
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateOrUpdateIngress(c *kubernetes.Clientset, ingress *extensions.Ingress) (*Ingress, error) {
	old, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Get(ingress.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			kubeIngress, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Create(ingress)
			if err != nil {
				return nil, err
			}
			return toIngress(kubeIngress), nil
		}
		return nil, err
	}
	ingress.Spec.DeepCopyInto(&old.Spec)
	kubeIngress, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Update(old)
	if err != nil {
		return nil, err
	}
	return toIngress(kubeIngress), nil
}

func GetIngressDetail(c *kubernetes.Clientset, name, namespace string) (*Ingress, error) {
	ingress, err := c.ExtensionsV1beta1().Ingresses(namespace).Get(name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return toIngress(ingress), nil
}

func DeleteIngress(c *kubernetes.Clientset, name, namespace string) error {
	return c.ExtensionsV1beta1().Ingresses(namespace).Delete(name, &metaV1.DeleteOptions{})
}
