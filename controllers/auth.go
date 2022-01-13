/*
 * @Description:认证相关
 * @Author: neozhang
 * @Date: 2022-01-03 16:39:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 23:02:23
 */
package controllers

import (
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"net/http"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

//Login 认证登录
func (c *AuthController) Login() {
	form := &forms.LoginForm{}
	errs := errors.New()
	//Get请求直接加载页面
	if c.Ctx.Input.IsPost() {
		//获取用户提交数据
		c.GetString("name")
		c.GetString("password")
		if err := c.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)
			if user == nil {
				//用户不存在
				errs.Add("default", "用户名或密码错误")
			} else if user.ValidPassword(form.Password) {
				//用户密码正确
				//记录用户状态
				sessionKey := beego.AppConfig.DefaultString("auth:Session", "") //TODO
				action := beego.AppConfig.DefaultString("auth:HomeAction", "")  //TODO

				c.SetSession(sessionKey, user.ID)
				c.Redirect(beego.URLFor(action), http.StatusFound)
			} else {
				//用户密码不正确
				errs.Add("default", "用户名或密码错误")
			}
		} else {
			errs.Add("default", "用户名或密码错误")
			//验证成功
			//验证失败
		}
	}
	c.Data["form"] = form
	c.Data["errors"] = errs
	//定义加载页面
	c.TplName = "auth/login.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	action := beego.AppConfig.DefaultString("auth:LogoutAction", "AuthCron") //TODO
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
