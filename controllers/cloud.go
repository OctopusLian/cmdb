package controllers

import (
	"strings"

	"cmdb/cloud"
	"cmdb/controllers/auth"
	"cmdb/forms"
	"cmdb/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type CloudPageController struct {
	LayoutController
}

func (c *CloudPageController) Index() {
	c.Data["expand"] = "asset_management"
	c.Data["menu"] = "cloud"
}

type CloudController struct {
	auth.LoginRequiredController
}

func (c *CloudController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("platform")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", query)
		queryCond = queryCond.Or("remark__icontains", query)
		cond = cond.AndCond(queryCond)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var platforms []*models.Platform
	qs.SetCond(cond).Limit(length).Offset(start).All(&platforms)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          platforms,
	}
	c.ServeJSON()
}

func (c *CloudController) Create() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.PlatformCreateForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				platform := &models.Platform{
					Name:       form.Name,
					Type:       form.Type,
					Addr:       form.Addr,
					Region:     form.Region,
					Key:        form.Key,
					Secrect:    form.Secrect,
					Remark:     form.Remark,
					CreateUser: c.User.Id,
				}
				if _, err := ormer.Insert(platform); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": platform,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "cloud/create.html"
		c.Data["types"] = models.PlatformTypes
	}
}

func (c *CloudController) Disable() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			platform := &models.Platform{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(platform) == nil {
				platform.Enable(false)
				ormer.Update(platform, "Status")
				json = map[string]interface{}{
					"code": 200,
					"text": "禁用成功",
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *CloudController) Enable() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			platform := &models.Platform{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(platform) == nil {
				platform.Enable(true)
				ormer.Update(platform, "Status")
				json = map[string]interface{}{
					"code": 200,
					"text": "启用成功",
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = map[string]interface{}{
		"code": 405,
		"text": "提交方式错误",
	}
	c.ServeJSON()
}

func (c *CloudController) Detail() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	if pk, err := c.GetInt("pk"); err == nil {
		platform := &models.Platform{Id: pk}
		ormer := orm.NewOrm()
		if ormer.Read(platform) == nil {
			json = map[string]interface{}{
				"code":   200,
				"text":   "获取成功",
				"result": platform,
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *CloudController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.PlatformModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				platform := &models.Platform{Id: form.Id}
				if ormer.Read(platform) == nil {
					platform.Name = form.Name
					platform.Type = form.Type
					platform.Addr = form.Addr
					platform.Region = form.Region
					if form.Key != "" {
						platform.Key = form.Key
					}
					if form.Secrect != "" {
						platform.Secrect = form.Secrect
					}

					platform.Remark = form.Remark

					if _, err := ormer.Update(platform); err == nil {
						json = map[string]interface{}{
							"code":   200,
							"text":   "更新成功",
							"result": platform,
						}
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		platform := &models.Platform{Id: -1}
		if pk, err := c.GetInt("pk"); err == nil {
			platform.Id = pk
			orm.NewOrm().Read(platform)
		}
		c.TplName = "cloud/modify.html"
		c.Data["platform"] = platform
		c.Data["types"] = models.PlatformTypes
	}
}

func (c *CloudController) Delete() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			platform := &models.Platform{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(platform) == nil {
				platform.Delete()
				if _, err := ormer.Update(platform, "delete_time"); err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "删除成功",
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

type VirtualMachinePageController struct {
	LayoutController
}

func (c *VirtualMachinePageController) Index() {
	c.Data["expand"] = "asset_management"
	c.Data["menu"] = "virtual_machine"
}

type VirtualMachineController struct {
	auth.LoginRequiredController
}

func (c *VirtualMachineController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	platformId, err := c.GetInt("platform")
	if err != nil {
		platformId = -1
	}

	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	ormer := orm.NewOrm()
	qs := ormer.QueryTable("virtual_machine")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total

	platform := &models.Platform{Id: platformId}

	if err := ormer.Read(platform); err == nil {
		cond = cond.And("platform__exact", platform)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	if query != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", query)
		queryCond = queryCond.Or("os__icontains", query)
		queryCond = queryCond.Or("public_addrs__icontains", query)
		queryCond = queryCond.Or("private_addrs__icontains", query)
		cond = cond.AndCond(queryCond)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var vms []*models.VirtualMachine
	qs.SetCond(cond).Limit(length).Offset(start).RelatedSel().All(&vms)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          vms,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Start() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			vm := &models.VirtualMachine{}
			ormer := orm.NewOrm()
			if ormer.QueryTable(vm).Filter("Id__exact", pk).RelatedSel().One(vm) == nil {
				if sdk, err := cloud.DefaultManager.Cloud(vm.Platform.Type); err == nil {
					sdk.Init(vm.Platform.Addr, vm.Platform.Key, vm.Platform.Secrect, vm.Platform.Region)
					sdk.StartInstance(vm.Key)
				}

				json = map[string]interface{}{
					"code":   200,
					"text":   "启动成功",
					"result": vm,
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *VirtualMachineController) Stop() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			vm := &models.VirtualMachine{}
			ormer := orm.NewOrm()
			if ormer.QueryTable(vm).Filter("Id__exact", pk).RelatedSel().One(vm) == nil {
				if sdk, err := cloud.DefaultManager.Cloud(vm.Platform.Type); err == nil {
					sdk.Init(vm.Platform.Addr, vm.Platform.Key, vm.Platform.Secrect, vm.Platform.Region)
					sdk.StopInstance(vm.Key)
				}

				json = map[string]interface{}{
					"code":   200,
					"text":   "停止成功",
					"result": vm,
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *VirtualMachineController) Restart() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
	}
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
		}
		if pk, err := c.GetInt("pk"); err == nil {
			vm := &models.VirtualMachine{}
			ormer := orm.NewOrm()
			if ormer.QueryTable(vm).Filter("Id__exact", pk).RelatedSel().One(vm) == nil {
				if sdk, err := cloud.DefaultManager.Cloud(vm.Platform.Type); err == nil {
					sdk.Init(vm.Platform.Addr, vm.Platform.Key, vm.Platform.Secrect, vm.Platform.Region)
					sdk.RestartInstance(vm.Key)
				}

				json = map[string]interface{}{
					"code":   200,
					"text":   "重启成功",
					"result": vm,
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}
