/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 19:05:58
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 19:05:59
 */
package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
