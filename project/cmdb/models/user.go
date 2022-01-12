/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 23:10:28
 */
package models

import (
	"cmdb/utils"
	"time"
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
	}

	return "" //TODO
}

//通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{}

	//TODO:issue 2
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		return user
	}
	return nil
}

// 查询用户
func QueryUser(q string) []*User {
	users := make([]*User, 0, 10)
	//TODO: panic
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return users
	}
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Name); err == nil {
			users = append(users, user)
		}
	}
	return users
}
