/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 22:50:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 23:18:38
 */
package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
)

type PasswordController struct {
	auth.AuthorizationController
}

func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{} //TODO
	errs := errors.New()
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证
			c.LoginUser.ValidPassword(form.Password); !ok {
				errs.Add("default","旧密码错误")
			} else if form.Password != form.Password2 {
				errs.Add("default","两次密码不一致")
			} else if form.OldPassword == form.Password {
				errs.Add("default","新旧密码不能一致")
			}
		}
	}
	c.TplName = "password/modify.html"
	c.Data["errors"] = errs
}
