package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Create a new config
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Create a new client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new informer factory
	factory := informers.NewSharedInformerFactory(client, 0)

	// Create a new event informer
	eventInformer := factory.Core().V1().Events().Informer()

	// Add event handler
	eventInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event, ok := obj.(metav1.Object)
			if !ok {
				log.Fatal("Failed to convert object to metav1.Object")
			}
			err := sendMessageToTelegram(fmt.Sprintf("Event Added: %s %s %s", event.GetNamespace(), event.GetName(), event.GetSelfLink()))
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			} else {
				log.Println("Message sent successfully!")
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			event, ok := newObj.(metav1.Object)
			if !ok {
				log.Fatal("Failed to convert object to metav1.Object")
			}
			err := sendMessageToTelegram(fmt.Sprintf("Event Added: %s %s %s", event.GetNamespace(), event.GetName(), event.GetSelfLink()))
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			} else {
				log.Println("Message sent successfully!")
			}
		},
		DeleteFunc: func(obj interface{}) {
			event, ok := obj.(metav1.Object)
			if !ok {
				log.Fatal("Failed to convert object to metav1.Object")
			}
			err := sendMessageToTelegram(fmt.Sprintf("Event Added: %s %s %s", event.GetNamespace(), event.GetName(), event.GetSelfLink()))
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			} else {
				log.Println("Message sent successfully!")
			}
		},
	})

	// Start the informer
	stopper := make(chan struct{})
	defer close(stopper)
	factory.Start(stopper)

	// Wait for the informer to sync
	if !cache.WaitForCacheSync(stopper, eventInformer.HasSynced) {
		log.Fatal("Failed to sync informer")
	}

	// Run indefinitely
	select {}
}
