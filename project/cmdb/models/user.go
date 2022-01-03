/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 22:24:15
 */
package models

import "cmdb/utils"

const (
	sqlQueryByName = "select id,name,password from user where name=?"
)

//User用户对象
type User struct {
	ID         int
	StaffID    string
	Name       string
	NickName   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
}

//验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return u.Password == utils.Md5Text(password)
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
