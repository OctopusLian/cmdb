/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 17:06:22
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 17:30:32
 */
package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
}
