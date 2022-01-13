/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:00:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 23:17:48
 */
package forms

//登录表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type PasswordModifyForm struct {
	Password    string
	Password2   string
	OldPassword string
}
