/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 17:06:22
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-04 17:22:33
 */
package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
}
