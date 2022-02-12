/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 21:45:27
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 17:28:11
 */
package services

import (
	"cmdb/forms"
	"cmdb/models"

	"github.com/astaxie/beego/orm"
)

type userService struct {
}

//通过用户ID获取用户信息
func (c *userService) GetByPk(pk int) *models.User {
	user := &models.User{
		Id: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}

	return nil
}

//通过用户名获取用户
func (c *userService) GetByName(name string) *models.User {
	user := &models.User{
		Name: name,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}

	return nil
}

//查询用户
func (c *userService) Query(q string) []*models.User {
	var users []*models.User
	queryset := orm.NewOrm().QueryTable(&models.User{})

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("nickname__icontains", q)
		query = query.Or("tel__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("email__icontains", q)
		query = query.Or("department__icontains", q)
		queryset.SetCond(query)
	}
	queryset.All(&users)
	return users
}

//修改用户信息
func (c *userService) ModifyUser(form *forms.UserModifyForm) {
	if user := c.GetByPk(form.Id); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
	}
}

//删除用户
func (c *userService) DeleteUser(pk int) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{
		Id: pk,
	})
}

//修改用户密码
func (c *userService) ModifyUserPassword(pk int, password string) {
	if user := c.GetByPk(pk); user != nil {
		user.Password = password
		ormer := orm.NewOrm()
		ormer.Update(user, "Password")
	}
}

//用户操作服务
var UserService = new(userService)
