/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 19:05:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 23:23:10
 */
package controllers

import (
	"cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.AuthorizationController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
