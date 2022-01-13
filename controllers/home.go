/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 19:05:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 22:55:58
 */
package controllers

import (
	"cmdb/base/controllers/auth"
)

// type HomeController struct {
// 	beego.Controller
// }

type HomeController struct {
	auth.AuthorizationController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
