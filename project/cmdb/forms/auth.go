/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:00:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 18:00:05
 */
package forms

//登录表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
