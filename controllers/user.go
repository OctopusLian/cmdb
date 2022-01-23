package controllers

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"cmdb/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
)

type UserPageController struct {
	LayoutController
}

func (c *UserPageController) Index() {
	c.Data["expand"] = "system_management"
	c.Data["menu"] = "user_management"
}

type UserController struct {
	auth.LoginRequiredController
}

func (c *UserController) List() {
	query := strings.TrimSpace(c.GetString("q"))
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")

	qs := orm.NewOrm().QueryTable("user")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	total, _ := qs.SetCond(cond).Count()

	qtotal := total
	if query != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", query)
		queryCond = queryCond.Or("department__icontains", query)
		queryCond = queryCond.Or("tel__icontains", query)
		queryCond = queryCond.Or("email__icontains", query)
		queryCond = queryCond.Or("addr__icontains", query)
		queryCond = queryCond.Or("remark__icontains", query)
		queryCond = queryCond.Or("department__icontains", query)
		cond = cond.AndCond(queryCond)
		qtotal, _ = qs.SetCond(cond).Count()
	}

	var users []*models.User
	qs.SetCond(cond).Limit(length).Offset(start).All(&users)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "成功",
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": qtotal,
		"result":          users,
	}
	c.ServeJSON()
}

func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				user := &models.User{
					Name:       form.Name,
					Gender:     form.Gender,
					Birthday:   &form.Birthday,
					Department: form.Department,
					Tel:        form.Tel,
					Email:      form.Email,
					Addr:       form.Addr,
					Remark:     form.Remark,
				}
				user.SetPassword(form.Password)
				if _, err := ormer.Insert(user); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": user,
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
		c.TplName = "user/create.html"
	}

}

func (c *UserController) Lock() {
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
			user := &models.User{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(user) == nil {
				user.Lock(true)
				ormer.Update(user, "Status")
				json = map[string]interface{}{
					"code": 200,
					"text": "锁定成功",
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Unlock() {
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
			user := &models.User{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(user) == nil {
				user.Lock(false)
				ormer.Update(user, "Status")
				json = map[string]interface{}{
					"code": 200,
					"text": "解锁成功",
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

func (c *UserController) Detail() {
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
	}
	if pk, err := c.GetInt("pk"); err == nil {
		user := &models.User{Id: pk}
		ormer := orm.NewOrm()
		if ormer.Read(user) == nil {
			json = map[string]interface{}{
				"code":   200,
				"text":   "获取成功",
				"result": user,
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.UserModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()

				user := &models.User{Id: form.Id}
				if ormer.Read(user) == nil {
					user.Name = form.Name
					user.Gender = form.Gender
					user.Birthday = &form.Birthday
					user.Department = form.Department
					user.Tel = form.Tel
					user.Email = form.Email
					user.Addr = form.Addr
					user.Remark = form.Remark
					if _, err := ormer.Update(user); err == nil {
						json = map[string]interface{}{
							"code":   200,
							"text":   "更新成功",
							"result": user,
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
		user := &models.User{Id: -1}
		if pk, err := c.GetInt("pk"); err == nil {
			user.Id = pk
			orm.NewOrm().Read(user)
		}
		c.TplName = "user/modify.html"
		c.Data["user"] = user
	}
}

func (c *UserController) Delete() {
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
			user := &models.User{Id: pk}
			ormer := orm.NewOrm()
			if ormer.Read(user) == nil {
				user.Delete()
				if _, err := ormer.Update(user, "delete_time"); err == nil {
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

func (c *UserController) Password() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.UserPasswordForm{User: c.User}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				ormer := orm.NewOrm()
				c.User.SetPassword(form.Password)
				if _, err := ormer.Update(c.User, "Password"); err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "修改密码成功",
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
		c.TplName = "user/password.html"
	}
}

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) List() {
	qs := orm.NewOrm().QueryTable("token")

	cond := orm.NewCondition()
	cond = cond.And("delete_time__isnull", true)

	var tokens []*models.Token
	qs.Filter("delete_time__isnull", true).Filter("create_user__exact", c.User.Id).All(&tokens)

	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "成功",
		"result": tokens,
	}
	c.ServeJSON()
}

func (c *TokenController) Generate() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		token := &models.Token{CreateUser: c.User.Id}
		ormer := orm.NewOrm()
		valid := &validation.Validation{}
		if _, _, err := ormer.ReadOrCreate(token, "create_user"); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			token.AccessKey = utils.RandString(64)
			token.SecrectKey = utils.RandString(64)
			if _, err := ormer.Update(token); err == nil {
				json = map[string]interface{}{
					"code":   200,
					"text":   "生成Token成功",
					"result": token,
				}
			} else {
				json = map[string]interface{}{
					"code": 500,
					"text": "服务器错误",
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		qs := orm.NewOrm().QueryTable("token")

		cond := orm.NewCondition()
		cond = cond.And("delete_time__isnull", true)

		var tokens []*models.Token
		qs.Filter("delete_time__isnull", true).Filter("create_user__exact", c.User.Id).All(&tokens)
		c.Data["tokens"] = tokens
		c.TplName = "token/index.html"
	}

}
