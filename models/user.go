/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 22:59:33
 */
package models

import (
	"cmdb/utils"
	"time"

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

func init() {
	orm.RegisterModel(new(User))
}
