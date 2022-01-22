/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-04 17:11:42
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-22 21:04:52
 */
package controllers

import (
	"cmdb/base/controllers/auth"
)

// UserController 用户控制管理器
type UserController struct {
	auth.AuthorizationController
}

// Query 查询用户
func (c *UserController) Query() {
	q := c.GetString("q")

	c.Data["users"] = services.userServices.Query(q) //TODO
	c.TplName = "user/query.html"
}

//修改用户
func (c *UserController) Modify() {

}

//删除用户
func (c *UserController) Delete() {

}
