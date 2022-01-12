/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-04 17:11:42
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-04 17:15:24
 */
package controllers

import (
	"cmdb/models"

	"github.com/astaxie/beego"
)

// UserController 用户控制管理器
type UserController struct {
	beego.Controller
}

// Query 查询用户
func (c *UserController) Query() {
	users := models.QueryUser("")
	c.Data["users"] = users
	c.TplName = "user/query.html"
}
