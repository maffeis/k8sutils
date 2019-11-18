package k8sutils

import (
	"flag"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// IClientSet the abstracts Kubernetes client API
type IClientSet interface {
	CoreV1() corev1.CoreV1Interface
}

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

// LoadSslCert loads an SSL certificate and its private key from a Secrets, and writes them to the
// pod file system. Those files are automatically deleted when the pod is destroyed.
func LoadSslCert(client IClientSet, nameSpace string, secretKey string, crtFile string, keyFile string) error {
	var secret *v1.Secret
	var error error

	secret, error = client.CoreV1().Secrets(nameSpace).Get(secretKey, metav1.GetOptions{})

	if error != nil {
		log.Fatalf("cannot retrieve secret: %s", error.Error())
	} else {
		caCert := secret.Data["ca.crt"]
		tlsCert := secret.Data["tls.crt"]
		privKey := secret.Data["tls.key"]

		if len(caCert) > 0 {
			if error := ioutil.WriteFile(crtFile, caCert, 0600); error != nil {
				log.Warnf("cannot write %s: %s", crtFile, error)
				return error
			}
		}

		if error := ioutil.WriteFile(crtFile, tlsCert, 0600); error != nil {
			log.Warnf("cannot write %s: %s", crtFile, error)
			return error
		}

		if error := ioutil.WriteFile(keyFile, privKey, 0600); error != nil {
			log.Warnf("cannot write %s: %s", keyFile, error)
			return error
		}
	}

	return nil
}
