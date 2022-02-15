/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:45:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-13 20:58:15
 */
package services

import (
	"cmdb/forms"
	"cmdb/models"
	"time"

	"github.com/astaxie/beego/orm"
)

type nodeService struct {
}

func (s *nodeService) Query(q string) []*models.Node {
	var nodes []*models.Node
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("hostname__icontains", q)
		qcond = qcond.Or("addr__icontains", q)
		cond = cond.AndCond(qcond)
	}
	queryset.SetCond(cond).All(&nodes)
	return nodes
}

func (s *nodeService) GetByPk(pk int) *models.Node {
	node := &models.Node{
		ID: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(node); err == nil {
		return node
	}
	return nil
}

func (s *nodeService) Delete(pk int) {
	if node := s.GetByPk(pk); node != nil {
		now := time.Now()
		node.DeletedAt = &now
		orm.NewOrm().Update(node, "DeletedAt")
	}
}

type jobService struct {
}

func (s *jobService) Query(q string) []*models.Job {
	var jobs []*models.Job
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = cond.Or("key__icontains", q)
		qcond = cond.Or("remark__icontains", q)
		cond = cond.AndCond(qcond)
	}
	queryset.SetCond(cond).All(&jobs)
	return jobs
}

func (s *jobService) GetByPk(pk int) *models.Job {
	job := &models.Job{
		ID: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(job); err == nil {
		return job
	}
	return nil
}

func (s *jobService) Delete(pk int) {
	if job := s.GetByPk(pk); job != nil {
		now := time.Now()
		job.DeletedAt = &now
		orm.NewOrm().Update(job, "DeletedAt")
	}
}

func (s *jobService) Create(form *forms.JobCreateForm) *models.Job {
	job := &models.Job{
		Key:    form.Key,
		Remark: form.Remark,
		Node:   NodeService.GetByPk(form.Node),
	}
	if _, err := orm.NewOrm().Insert(job); err == nil {
		return job
	}
	return nil
}

type targetService struct {
}

func (s *targetService) Query(q string) []*models.Target {
	var targets []*models.Target
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = cond.Or("name__icontains", q)
		qcond = cond.Or("remark__icontains", q)
		qcond = cond.Or("addr__icontains", q)
		cond.AndCond(qcond)
	}
	queryset.SetCond(cond).All(&targets)
	return targets
}

func (s *targetService) GetByPk(pk int) *models.Target {
	target := &models.Target{
		ID: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(target); err == nil {
		return target
	}
	return nil
}

func (s *targetService) Delete(pk int) {
	if target := s.GetByPk(pk); target != nil {
		now := time.Now()
		target.DeletedAt = &now
		orm.NewOrm().Update(target, "DeletedAt")
	}
}

var (
	NodeService   = new(nodeService)
	JobService    = new(jobService)
	TargetService = new(targetService)
)
