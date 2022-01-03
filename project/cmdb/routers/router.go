/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 17:06:22
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 22:14:28
 */
package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
}
