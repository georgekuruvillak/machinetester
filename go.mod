module github.com/georgekuruvillak/machinetester

go 1.13

require (
	github.com/gardener/machine-controller-manager v0.26.3
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.5.0
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90 // kubernetes-1.16.0
