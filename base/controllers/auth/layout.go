/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-12 15:29:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 20:20:22
 */
package auth

import (
	"strings"

	"github.com/astaxie/beego"
)

//布局控制器基础
type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) GetNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()
	c.Layout = "base/layouts/layout.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyle"] = ""
	c.LayoutSections["SectionScript"] = ""

	c.Data["nav"] = c.GetNav()
	c.Data["subnav"] = ""
	c.Data["title"] = beego.AppConfig.DefaultString("AppName", "CMDB")
}
