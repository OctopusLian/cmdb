/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:15:34
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-15 21:11:08
 */
package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/services"
	"fmt"
	"net/http"

	"cmdb/forms"

	"github.com/astaxie/beego"
)

type prometheusController struct {
	auth.LayoutController
}

func (c *prometheusController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "prometheus"
	c.Data["subnav"] = c.GetNav()
}

type NodeController struct {
	prometheusController
}

func (c *NodeController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "prometheus"
	c.Data["subnav"] = "node"
}

func (c *NodeController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["nodes"] = services.NodeService.Query(q)
	c.Data["q"] = q
	c.TplName = "prometheus/node/query.html"
}

func (c *NodeController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.NodeService.Delete(pk)
	}
	c.Redirect(beego.URLFor("NodeController.Query"), http.StatusFound)
}

type JobController struct {
	prometheusController
}

func (c *JobController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "prometheus"
	c.Data["subnav"] = "node"
}

func (c *JobController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["jobs"] = services.JobService.Query(q)
	c.Data["q"] = q

	c.TplName = "prometheus/job/query.html"
}

func (c *JobController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.JobService.Delete(pk)
	}
	c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
}

func (c *JobController) Create() {
	form := &forms.JobCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证
			services.JobService.Create(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.JobService.Create(form)
	c.TplName = "prometheus/job/create.html"
}

func (c *JobController) Modify() {
	form := &forms.JobCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证
			services
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	} else {
		if pk, err := c.GetInt("pk"); err == nil {
			job := services.JobService.GetByPk(pk)
			// form.ID = job.ID
			form.Key = job.Key
			form.Remark = job.Remark
			form.Node = job.Node.ID
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.JobService
	c.TplName = "prometheus/job/modify.html"
}

type TargetController struct {
	prometheusController
}

func (c *TargetController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "prometheus"
	c.Data["subnav"] = "node"
}

func (c *TargetController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["targets"] = services.TargetService
	c.Data["q"] = q

	c.TplName = "prometheus/target/query.html"
}

func (c *TargetController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.TargetService.Delete(pk)
	}
	c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
}

func (c *TargetController) Create() {
	form := &forms.JobCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证
			services.TargetService.
				c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services
	c.TplName = "prometheus/target/create.html"
}

func (c *TargetController) Modify() {
	//Get 查询值显示
	//POST 更新
	form := &forms.TargetModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			//验证
			services.TargetService.Modify()
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	} else {
		if pk, err := c.GetInt("pk"); err == nil {
			if target := services.TargetService.GetByPk(pk); target != nil {
				form = forms.NewTargetModifyForm(target)
			}
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.TargetService.Query("")
	c.TplName = "prometheus/target/modify.html"
}

type AlertController struct {
	prometheusController
}

func (c *AlertController) Query() {
	form := forms.NewAlertQueryParams(c.Input())
	if err := c.ParseForm(form); err == nil {
		fmt.Printf("%#v\n", form.PageQueryParams)
		fmt.Printf("%#v\n", form)
		c.Data["page"] = services
	}
	c.Data["form"] = form
	c.TplName = "prometheus/alert/query.html"
}
