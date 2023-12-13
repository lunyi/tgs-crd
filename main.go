package main

import (
	"context"
	"fmt"
	"log"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	v1 "srvptt2.online/pkg/apis/srvptt2.online/v1"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln(err)
	}

	config.APIPath = "/apis/"
	config.NegotiatedSerializer = v1.Codecs.WithoutConversion()
	config.GroupVersion = &v1.GroupVersion

	client, err := rest.RESTClientFor(config)

	if err != nil {
		log.Fatalln(err)
	}

	lobby := v1.Lobby{}
	err = client.Get().Namespace("dev").Resource("srvptt2.online").Name("lobby-t1").Do(context.TODO()).Into(&lobby)

	if err != nil {
		log.Fatalln(err)
	}

	newObj := lobby.DeepCopy()
	newObj.Spec.Name = "test2"

	fmt.Println(lobby.Spec.Name)
	fmt.Println(lobby.Spec.Image)
	fmt.Println(lobby.Spec.Domain)
	fmt.Println(lobby.Spec.Replicas)

	fmt.Println(newObj.Spec.Name)
}
