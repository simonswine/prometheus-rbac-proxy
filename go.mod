module github.com/simonswine/prometheus-rbac-proxy

go 1.12

require (
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.6.0
	github.com/onsi/gomega v1.4.2
	github.com/prometheus/prometheus v2.5.0+incompatible
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	sigs.k8s.io/controller-runtime v0.3.0
)
