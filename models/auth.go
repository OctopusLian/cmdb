/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-23 10:26:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:19:57
 */
package models

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

// type User struct {
// 	Model
// 	Id          int        `orm:"column(id);" json:"id"`
// 	Name        string     `orm:"column(name);size(32);" json:"name"`
// 	Password    string     `orm:"column(password);size(1024);" json:"-"`
// 	Gender      int        `orm:"column(gender);default(0);" json:"gender"`
// 	Birthday    *time.Time `orm:"column(birthday);type(date);null;default(null);" json:"birthday"`
// 	Tel         string     `orm:"column(tel);size(32);" json:"tel"`
// 	Email       string     `orm:"column(email);size(32);" json:"email"`
// 	Addr        string     `orm:"column(addr);size(1024);" json:"addr"`
// 	Remark      string     `orm:"column(remark);size(1024);" json:"remark"`
// 	Department  string     `orm:"column(department);size(32);" json:"department"`
// 	Status      int        `orm:"column(status);default(0);" json:"status"`
// 	IsSuperuser bool       `orm:"column(is_superuser);default(0);" json:"is_superuser"`
// }

// func (u *User) ValidatePassword(password string) bool {
// 	salt, _ := utils.SplitMd5Salt(u.Password)
// 	return u.Password == utils.Md5Salt(password, salt)
// }

// func (u *User) SetPassword(password string) {
// 	u.Password = utils.Md5Salt(password, "")
// }

// func (u *User) IsLock() bool {
// 	return u.Status == statusLock
// }

// func (u *User) Lock(lock bool) bool {
// 	if lock {
// 		u.Status = statusLock
// 	} else {
// 		u.Status = statusUnlock
// 	}
// 	return true
// }

// type Token struct {
// 	Model
// 	Id         int    `orm:"column(id);" json:"id" `
// 	AccessKey  string `orm:"column(access_key);size(512);"`
// 	SecrectKey string `orm:"column(secrect_key);size(512);"`
// 	Status     int    `orm:"column(status);default(0);"`
// 	CreateUser int    `orm:"column(create_user);default(0)`
// }

func (t *Token) ValidateSignature(signature string, input *context.BeegoInput) bool {
	return t.SecrectKey == signature
}

func init() {
	orm.RegisterModel(new(User), new(Token))
}
