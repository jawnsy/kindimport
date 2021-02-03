package kindimport

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/scheme"
)

var _ = scheme.ParameterCodec

func main() {
	testfunc(nil)
}

func testfunc(client kubernetes.Interface) {
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("abc").
		Namespace("abc").
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: "abc",
		Command:   nil,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
	}, scheme.ParameterCodec)

}
