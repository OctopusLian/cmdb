package routers

import (
	"cmdb/controllers/auth"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&auth.AuthController{})
}
