/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 22:50:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-23 14:07:42
 */
package controllers

//"regexp"

// type PasswordController struct {
// 	auth.AuthorizationController
// }

// func (c *PasswordController) Modify() {
// 	form := &forms.PasswordModifyForm{} //TODO
// 	errs := errors.New()
// 	text := ""
// 	if c.Ctx.Input.IsPost() {
// 		if err := c.ParseForm(form); err == nil {

// 			//验证
// 			if ok := c.LoginUser.ValidPassword(form.OldPassword); !ok {
// 				errs.Add("default", "旧密码错误")
// 			} else {
// 				valid := &validation.Validation{}
// 				if hasError, err := valid.Valid(form); err != nil {
// 					errs.Add("default", err.Error())
// 				} else if hasError {
// 					//errs.AddValidation(valid)
// 				} else {
// 					//models.ModifyUserPassword(c.LoginUser.ID, form.Password)
// 				}

// 			}

// 		}
// 	}
// 	c.TplName = "password/modify.html"
// 	c.Data["errors"] = errs
// 	c.Data["text"] = text
// }
