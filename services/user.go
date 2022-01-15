/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-14 13:37:57
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 23:15:38
 */
package services

import (
	"cmdb/forms"
	"cmdb/models"

	"github.com/astaxie/beego/orm"
)

type userService struct {
}

//用户操作服务
var userServices = new(userService)

//通过用户ID获取用户信息
func (s *userService) GetByPk(pk int) *models.User {
	user := &models.User{
		ID: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}

	return nil
}

//通过用户名获取用户
func (s *userService) GetByName(name string) *models.User {
	user := &models.User{
		Name: name,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 查询用户
func (s *userService) Query(q string) []*models.User {
	var users []*models.User
	queryset := orm.NewOrm().QueryTable(&models.User{})
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__iccontains", q)
		cond = cond.Or("nickname__iccontains", q)
		cond = cond.Or("tel__iccontains", q)
		cond = cond.Or("addr__iccontains", q)
		cond = cond.Or("email__iccontains", q)
		cond = cond.Or("department__iccontains", q)
		//queryset = queryset.SetCond((cond)
	}
	queryset.All(&users)
	return users
}

//修改用户信息
func (s *userService) Modify(form *forms.UserModifyForm) {
	if user := s.GetByPk(form.ID); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
	}
}

//删除用户
func (s *userService) Delete(pk int) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{ID: pk})
}

//修改用户密码
func (s *userService) ModifyUserPassword(pk int, password string) {
	if user := s.GetByPk(pk); user != nil {
		user.Password = password
		ormer := orm.NewOrm()
		ormer.Update(user, "Password")
	} else {
		//记录错误
	}
}
