/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:17:09
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 21:37:44
 */
package routers

import (
	"github.com/astaxie/beego"

	"cmdb/controllers"
	v1 "cmdb/controllers/api/v1"
	v2 "cmdb/controllers/api/v2"
	"cmdb/controllers/auth"
)

func init() {
	//beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/home", &controllers.HomeController{})

	// 认证
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	// 认证
	beego.AutoRouter(&auth.AuthController{})

	// Dashboard
	beego.AutoRouter(&controllers.DashboardPageController{})

	// Dashboard
	beego.AutoRouter(&controllers.DashboardController{})

	// 用户页面
	beego.AutoRouter(&controllers.UserPageController{})

	// 用户
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
	//beego.AutoRouter(&controllers.PasswordController{})

	// 云平台页面
	beego.AutoRouter(&controllers.CloudPlatformPageController{})

	// 云平台
	beego.AutoRouter(&controllers.CloudPlatformController{})

	// 云主机页面
	beego.AutoRouter(&controllers.VirtualMachinePageController{})

	// 云主机
	beego.AutoRouter(&controllers.VirtualMachineController{})

	// 云主机页面
	beego.AutoRouter(&controllers.AgentPageController{})

	// 云主机
	beego.AutoRouter(&controllers.AgentController{})

	// 资源使用率页面
	beego.AutoRouter(&controllers.ResourcePageController{})

	// 资源使用率
	beego.AutoRouter(&controllers.ResourceController{})

	// 告警页面
	beego.AutoRouter(&controllers.AlarmPageController{})

	// 告警
	beego.AutoRouter(&controllers.AlarmController{})

	//prometheus
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.TargetController{})

	v1Namespace := beego.NewNamespace("/v1",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
		beego.NSRouter("api/task/:uuid/", &v1.APIController{}, "*:Task"),
		beego.NSRouter("api/result/:uuid/", &v1.APIController{}, "*:TaskResult"),
	)

	beego.AddNamespace(v1Namespace)

	v2Namespace := beego.NewNamespace("/v2",
		beego.NSRouter("api/heartbeat/:uuid/", &v2.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v2.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v2.APIController{}, "*:Log"),
		beego.NSRouter("api/task/:uuid/", &v2.APIController{}, "*:Task"),
		beego.NSRouter("api/result/:uuid/", &v2.APIController{}, "*:TaskResult"),
	)
	beego.AddNamespace(v2Namespace)
}
