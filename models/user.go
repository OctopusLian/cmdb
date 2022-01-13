/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-13 23:15:27
 */
package models

import (
	"cmdb/utils"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/beego/beego/v2/adapter/orm"
)

const (
	sqlQueryByName = "select id,name,password from user where name=?"
	sqlQuery       = "select id from user"
)

//User用户对象
type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	NickName   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
}

//验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return u.Password == utils.Md5Text(password)
}

//性别显示
func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}

//状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}

func GetUserByPk(pk int) *User {
	user := &User{
		ID: pk,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}

	return nil
}

//通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{
		Name: name,
	}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 查询用户
func QueryUser(q string) []*User {
	var users []*User
	queryset := orm.NewOrm().QueryTable(&User{})
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__iccontains", q)
		cond = cond.Or("nickname__iccontains", q)
		cond = cond.Or("tel__iccontains", q)
		cond = cond.Or("addr__iccontains", q)
		cond = cond.Or("email__iccontains", q)
		cond = cond.Or("department__iccontains", q)
		queryset = queryset.SetCond((cond)
	}
	queryset.All(&useers)
	return users
}

//修改用户信息
//func ModifyUser(form *forms.User)

//删除用户
func DeleteUser(pk int) {
	ormer := orm.NewOrm()
	ormer.Delete(&User{ID:pk})
}

func ModifyUserPassword(pk int,password string) {
	if user := GetUserByPk(pk); user != nil {
		user.Password = password
		ormer := orm.NewOrm()
		ormer.Update(user,"Password")
	}
}

func init() {
	orm.RegisterModel(new(User))
}