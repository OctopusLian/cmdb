/*
 * @Description:认证相关
 * @Author: neozhang
 * @Date: 2022-01-03 16:39:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 16:52:58
 */
package controllers

import "github.com/astaxie/beego"

type AuthController struct {
	beego.Controller
}

//Login 认证登录
func (c *AuthController) Login() {
	//Get请求直接加载页面
	c.Ctx.Input.IsPost()
	//Post请求进行验证
	//验证成功
	//验证失败

	//定义加载页面
	c.TplName = "auth/login.html"
}
