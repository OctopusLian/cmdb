package base

import (
	"fmt"
	"strings"

	"cmdb/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	controller, action := c.GetControllerAndAction()
	controller, action = strings.TrimSuffix(utils.Snake(controller), "_controller"), utils.Snake(action)
	c.TplName = fmt.Sprintf("%s/%s.%s", controller, action, "html")
}
