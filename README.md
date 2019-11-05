# Go module `k8sutils`

Module `github.com/maffeis/k8sutils` provides simple utility functions for developing Kubernetes applications in go.
The following example checks whether the go application is running on Kubernetes. In case it is, the available Kubernetes nodes are listed.

```go
import (
	"fmt"

	log "github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"vipera.com/k8sutils"
)

func main() {
	if k8sutils.IsRunningOnKubernetes() {
		var client *kubernetes.Clientset
		var error error

		client, error = k8sutils.KubernetesConfig()

		if error != nil {
			panic(error)
		} else {
			var ver *version.Info

			ver, error = client.ServerVersion()

			log.Infof("running on Kubernetes: %s", ver.String())

			list, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
			if err != nil {
				log.Fatal("cannot retrieve Kubernetes nodes: %s", err.Error())
			}

			for _, node := range list.Items {
				fmt.Printf("Node: %s\n", node.Name)
			}
		}
	} else {
		log.Infof("not running on Kubernetes")
	}
}
```