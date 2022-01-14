/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-13 13:20:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-14 13:21:01
 */
package forms

type UserModifyForm struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
}
