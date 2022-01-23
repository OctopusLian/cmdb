package routers

import (
	"github.com/astaxie/beego"

	"cmdb/controllers"
)

func init() {

	// home
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	// 云平台
	beego.AutoRouter(&controllers.CloudPageController{})
	beego.AutoRouter(&controllers.CloudController{})

	// 虚拟机
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})

	// 用户
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
}
