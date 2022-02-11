/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:45:41
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 21:59:25
 */
package services

import (
	"github.com/astaxie/beego"
	"github.com/client-go/kubernetes"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type deploymentService struct{}

func (s *deploymentService) Query() {
	path := beego.AppConfig.DefaultString("k8s::path", "etc/k8s/kube.conf")
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return []appsV1.deployment{}
	}
	kubernetes.NewForConfig
}

var DeploymentService = new(deploymentService)
