/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:38:59
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 19:00:21
 */
package models

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
	return false
}

//通过用户名获取用户
func GetUserByName(name string) *User {
	return nil
}
