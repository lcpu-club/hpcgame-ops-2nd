package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
)

const (
	configMapCount = 1000
	configMapSize  = 1 * 1024 * 1024 // 1 MiB
	namespace      = "test-ns"
)

func main() {
	// Load kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", "/home/ubuntu/.kube/config")
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	config.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(50, 100)

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start time
	start := time.Now()

	// Generate ConfigMaps in parallel
	for i := 0; i < configMapCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			createConfigMap(clientset, i)
		}(i)
	}

	// Wait for all ConfigMaps to be created
	wg.Wait()

	// Calculate and print the total time taken
	elapsed := time.Since(start)
	fmt.Printf("Created %d ConfigMaps in %s\n", configMapCount, elapsed)
}

func createConfigMap(clientset *kubernetes.Clientset, index int) {
	// Generate a unique name for the ConfigMap
	name := fmt.Sprintf("configmap-%d-%d", 3, index)

	// Generate 1 MiB of data
	data := make(map[string]string)
	data["data"] = generateRandomString(configMapSize)

	// Create the ConfigMap
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}

	// Create the ConfigMap in the cluster
	_, err := clientset.CoreV1().ConfigMaps(namespace).Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating ConfigMap %s: %v", name, err)
	} else {
		log.Printf("Created ConfigMap %s", name)
	}
}

func generateRandomString(size int) string {
	// Generate a random string of the given size
	// This is a simple implementation, you might want to use a more efficient method
	b := make([]byte, size)
	for i := range b {
		b[i] = 'a' + byte(i%26)
	}
	return string(b)
}
