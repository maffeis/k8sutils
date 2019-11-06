module github.com/maffeis/k8sutils

go 1.13

require (
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	k8s.io/apimachinery v0.0.0-20191025225532-af6325b3a843
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
