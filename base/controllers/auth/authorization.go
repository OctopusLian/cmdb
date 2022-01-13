/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 22:52:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-12 23:12:18
 */
package auth

import (
	"cmdb/models"

	"github.com/astaxie/beego"
)

type AuthorizationController struct {
	//TODO
	beego.Controller
	LoginUser models.User
}
