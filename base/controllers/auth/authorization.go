/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 22:52:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 18:14:10
 */
package auth

import (
	"cmdb/models"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
)

//所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	beego.Controller
	LoginUser *models.User
}

//用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionValue := c.GetSession(sessionKey)
	c.Data["loginUser"] = nil

	if sessionValue != nil {
		if pk, ok := sessionValue.(int); ok {
			if user := services.UserService.GetByPk(pk); user != nil {
				c.Data["loginUser"] = user
				c.LoginUser = user
				return
			}
		}
	}

	action := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
