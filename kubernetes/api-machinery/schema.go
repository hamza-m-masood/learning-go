package api_machinery

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	//create a scheme
	Scheme := runtime.NewScheme()
	//Add pod and configmap type in scheme
	pod := corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "test",
					Image: "nginx",
				},
			},
		},
	}
	Scheme.AddKnownTypes(schema.GroupVersion{
		Group:   "",
		Version: "v1",
	}, &pod, &corev1.ConfigMap{
		Data: map[string]string{
			"test": "data",
		},
	})

	//get map of known types
	kt := Scheme.KnownTypes(schema.GroupVersion{
		Group:   "",
		Version: "v1",
	})
	fmt.Println(kt)

	//get group-version of known types
	gv := Scheme.VersionsForGroupKind(schema.GroupKind{
		Group: "",
		Kind:  "ConfigMap",
	})
	fmt.Println(gv)

	//create new object
	obj, _ := Scheme.New(schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Pod",
	})
	fmt.Println(obj)

	//get objectkind of known types
	ok, _, _ := Scheme.ObjectKinds(obj)
	fmt.Println(ok)

}
