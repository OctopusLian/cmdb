/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 20:49:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 21:07:29
 */
package main

import (
	"log"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	configPath := "etc/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	namesapce := "default"
	var replicas int32 = 3
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				"app": "nginx",
				"env": "dev",
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
					"env": "dev",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
						"env": "dev",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.16.1",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.Protocol(coreV1.IPv4Protocol), //TCP
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

}
