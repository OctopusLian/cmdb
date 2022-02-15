/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-13 20:41:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-15 21:09:10
 */
package forms

import "cmdb/models"

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

type TargetCreateForm struct {
	Name   string `form:"key"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}

type TargetModifyForm struct {
	ID     int    `form:"id"`
	Name   string `form:"key"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}

func NewTargetModifyForm(target *models.Target) *TargetModifyForm {
	form := &TargetModifyForm{}
	form.ID = target.ID
	form.Name = target.Name
	form.Remark = target.Remark
	form.Addr = target.Addr
	form.Job = target.Job.ID
	return form
}
