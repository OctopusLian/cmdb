/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:17:56
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 21:28:26
 */
package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type k8sController struct {
	LayoutController
}

func (c *k8sController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "k8s"
	c.Data["subnav"] = c.GetNav()
}

type DeploymentController struct {
	k8sController
}

func (c *DeploymentController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["deployment"] = services
	c.Data["q"] = q

	c.TplName = "k8s/deployment/query.html"
}

func (c *DeploymentController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {

	}
	c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
}
