/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:45:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 21:31:34
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

func (s *jobService) Query(q string) []*models.Job {
	var jobs []*models.Job
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	if q != nil {
		cond := orm.NewCondition()
		cond = cond.Or("key__icontains", q)
		cond = cond.Or("remark__icontains", q)
		queryset.SetCond(cond)
	}
	queryset.All(&jobs)
	return jobs
}

type targetService struct {
}

func (s *targetService) Query(q string) []*models.Target {
	var targets []*models.Target
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	if q != nil {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("remark__icontains", q)
		cond = cond.Or("addr__icontains", q)
		queryset.SetCond(cond)
	}
	queryset.All(&targets)
	return targets
}

var (
	NodeService   = new(nodeService)
	JobService    = new(jobService)
	TargetService = new(targetService)
)
