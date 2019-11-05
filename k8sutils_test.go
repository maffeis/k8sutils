package k8sutils

import (
	"testing"
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
