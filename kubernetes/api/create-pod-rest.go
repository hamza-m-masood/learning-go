package api

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"net/http"
)

func main() {
	if err := createPod(); err != nil {
		fmt.Println("Could not create pod: ", err)
	}
}

func createPod() error {
	fmt.Println("hello")
	pod := createPodObject()
	serializer := getJSONSerializer()
	postBody, err := serializePodObject(serializer, pod)
	if err != nil {
		return err
	}

	reqCreate, err := buildPostRequest(postBody)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(reqCreate)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode < 300 {
		createdPod, err := deserializePodBody(serializer, body)
		if err != nil {
			return err
		}
		json, err := stdjson.MarshalIndent(createdPod, "", " ")

		if err != nil {
			return err
		}
		fmt.Printf("%s\n", json)
	} else {
		status, err := deserializeStatusBody(serializer, body)
		if err != nil {
			return err
		}
		json, err := stdjson.MarshalIndent(status, "", " ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", json)
	}
	return nil
}

func createPodObject() *corev1.Pod {
	pod := corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "runtime",
					Image: "nginx",
				},
			},
		},
	}
	pod.SetName("my-pod")
	pod.SetLabels(map[string]string{
		"app.kubernetes.io/component": "my-component",
		"app.kubernetes.io/name":      "a-name",
	})
	return &pod
}

func getJSONSerializer() runtime.Serializer {
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(
		schema.GroupVersion{Group: "", Version: "v1"},
		&corev1.Pod{},
		&metav1.Status{},
	)
	return json.NewSerializerWithOptions(
		json.SimpleMetaFactory{},
		nil,
		scheme,
		json.SerializerOptions{},
	)
}

func serializePodObject(serializer runtime.Serializer, pod *corev1.Pod) (io.Reader, error) {
	var buf bytes.Buffer
	err := serializer.Encode(pod, &buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func buildPostRequest(body io.Reader) (*http.Request, error) {
	reqCreate, err := http.NewRequest("POST", "http://127.0.0.1:8001/api/v1/namespaces/default/pods", body)
	if err != nil {
		return nil, err
	}
	reqCreate.Header.Add("Accept", "application/json")
	reqCreate.Header.Add("Content-Type", "application/json")
	return reqCreate, nil
}
func deserializePodBody(serializer runtime.Serializer, body []byte) (*corev1.Pod, error) {
	var result corev1.Pod
	_, _, err := serializer.Decode(body, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func deserializeStatusBody(serializer runtime.Serializer, body []byte) (*metav1.Status, error) {
	var status metav1.Status
	_, _, err := serializer.Decode(body, nil, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
