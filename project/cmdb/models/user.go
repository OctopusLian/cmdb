/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-04 17:30:55
 */
package models

import "cmdb/utils"

const (
	sqlQueryByName = "select id,name,password from user where name=?"
	sqlQuery       = "select id from user"
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
