/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-13 20:41:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-13 20:43:25
 */
package forms

type JobCreateForm struct {
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type JobModifyForm struct {
	ID     int    `form:"id"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}
