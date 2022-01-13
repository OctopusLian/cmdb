/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 22:50:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-13 23:19:06
 */
package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/models"
	"regexp"

	"github.com/beego/beego/v2/core/validation"
	//"google.golang.org/protobuf/internal/encoding/text"
)

type PasswordController struct {
	auth.AuthorizationController
}

func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{} //TODO
	errs := errors.New()
	text := ""
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if ok := c.LoginUser.ValidPassword(form.OldPassword); !ok {
				errs.Add("default", "旧密码错误")
			} else {
				valid := &validation.Validation{}
				//验证 密码范围数字，大小写英文字母、特殊字符
				passwordRegax := "^[0-9a-zA-Z_.]" //TODO

				valid.Match(form.Password, regexp.MustCompile(passwordRegax), "default.default.default")
				if isMatch := regexp.MatchString(passwordRegax, form.Password); !isMatch {
					errs.Add("default", "密码只能由大写、小写、数字、特殊字符组成")
				} else if form.Password != form.Password2 {
					errs.Add("default", "两次密码不一致")
				} else if form.OldPassword == form.Password {
					errs.Add("default", "新旧密码不能一致")
				} else {
					models.ModifyUserPassword(c.LoginUser.ID, form.Password)
					text = "修改密码成功"
				}
			}

		}
	}
	c.TplName = "password/modify.html"
	c.Data["errors"] = errs
	c.Data["text"] = text
}
