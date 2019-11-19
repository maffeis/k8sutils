package k8sutils

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestIsRunningOnKubernetes(t *testing.T) {
	want := false
	if got := IsRunningOnKubernetes(); got != want {
		t.Errorf("IsRunningOnKubernetes() returned true")
	}
}

func TestIsRunningOnDocker(t *testing.T) {
	want := false
	if got := IsRunningOnDocker(); got != want {
		t.Errorf("IsRunningOnDocker() returned true")
	}
}

type MockClientset struct {
}

func (d MockClientset) CoreV1() ICoreV1Interface {
	return new(MockCoreV1Interface)
}

type MockCoreV1Interface struct {
}

type MockSecrets struct {
}

func (d MockSecrets) Get(name string, options metav1.GetOptions) (*v1.Secret, error) {
	return new(v1.Secret), nil
}

func (d MockSecrets) Create(*v1.Secret) (*v1.Secret, error) {
	return nil, nil
}

func (d MockSecrets) Delete(name string, options *metav1.DeleteOptions) error {
	return nil
}

func (d MockSecrets) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return nil
}

func (d MockSecrets) List(opts metav1.ListOptions) (*v1.SecretList, error) {
	return nil, nil
}

func (d MockSecrets) Update(*v1.Secret) (*v1.Secret, error) {
	return nil, nil
}

func (d MockSecrets) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return nil, nil
}

func (d MockSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Secret, err error) {
	return nil, nil
}

func (d MockCoreV1Interface) Secrets(namespace string) corev1.SecretInterface {
	return new(MockSecrets)
}

func TestLoadSslCert(t *testing.T) {
	LoadSslCert(new(MockClientset), "ns", "key", "crtFile", "keyFile")
}
