package k8sutils

import (
	"flag"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// IsRunningOnKubernetes checks whether the application is running on
// a Kubernetes cluster.
func IsRunningOnKubernetes() bool {
	if _, err := os.Stat("/var/run/secrets/kubernetes.io"); !os.IsNotExist(err) {
		return true
	}

	return false
}

// IsRunningOnDocker checks whether the application is running in
// a Docker container.
func IsRunningOnDocker() bool {
	if _, err := os.Stat("/.dockerenv"); !os.IsNotExist(err) {
		return true
	}

	return false
}

// KubernetesConfig configures the application for Kubernetes.
func KubernetesConfig() (*kubernetes.Clientset, error) {
	kubeconfig := ""
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()

	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	var (
		config *rest.Config
		err    error
	)

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	client := kubernetes.NewForConfigOrDie(config)

	return client, nil
}
