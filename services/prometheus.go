/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:45:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 17:42:41
 */
package services

import (
	"cmdb/models"

	"github.com/astaxie/beego/orm"
)

type nodeService struct {
}

func (s *nodeService) Query(q string) []*models.Node {
	var nodes []*models.Node
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	if q != nil {
		cond := orm.NewCondition()
		cond = cond.Or("hostname__icontains", q)
		cond = cond.Or("addr__icontains", q)
		queryset.SetCond(cond)
	}
	queryset.All(&nodes)
	return nodes
}

type jobService struct {
}

func (s *nodeService) Query(q string) []*models.Job {
	var jobs []*models.Jobs
}
