/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-14 13:21:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-14 13:22:38
 */
package forms

import (
	"regexp"

	"github.com/astaxie/beego/validation"
)

type PasswordModifyForm struct {
	OldPassword string `form:"old_password"`
	Password    string `form:"password"`
	Password2   string `form:"password2"`
}

func (f *PasswordModifyForm) Valid(validation *validation.Validation) {
	//验证 密码范围数字，大小写英文字母、特殊字符
	passwordRegax := "^[0-9a-zA-Z_.]" //TODO

	validation.Match(f.Password, regexp.MustCompile(passwordRegax), "default.default.default")
	if validation.HasErrors() {
		return
	} else if f.Password != f.Password2 {
		// if isMatch, _ := regexp.MatchString(passwordRegax, form.Password); !isMatch {
		// 	errs.Add("default", "密码只能由大写、小写、数字、特殊字符组成")

		//errs.Add("default", "两次密码不一致")
		validation.AddError("default.default", "两次密码不一致")
	} else if f.OldPassword == f.Password {
		//errs.Add("default", "新旧密码不能一致")
		validation.AddError("default.default", "新旧密码不能一致")
	}
}
