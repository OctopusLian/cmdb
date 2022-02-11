/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 20:29:02
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 20:48:29
 */
package main

import (
	"context"
	"fmt"
	"log"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func List() {
	configPath := "etc/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodeList.Items {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s,\n",
			&node.Name,
			&node.Status.Phase,
			&node.Status.Addresses,
			&node.Status.NodeInfo.OSImage,
			&node.Status.NodeInfo.KernelVersion,
			&node.Status.NodeInfo.OperatingSystem,
			&node.Status.NodeInfo.Architecture,
			node.CreationTimestamp,
		)
	}

	namespaceList, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	for _, namespace := range namespaceList.Items {
		fmt.Println(namespace.Name, namespace.CreationTimestamp)
	}

	servicesList, _ := clientset.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
	for _, service := range servicesList.Items {
		fmt.Println(service.Name)
	}

	deploymentList, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metaV1.ListOptions{})
	for _, deployment := range deploymentList.Items {
		fmt.Println(deployment.Name)
	}
}
