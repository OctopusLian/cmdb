package controllers

import (
	"fmt"
	"strings"

	"cmdb/controllers/auth"
	"cmdb/utils"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare() {
	c.LoginRequiredController.Prepare()
	if c.User != nil {
		controller, action := c.GetControllerAndAction()
		controller, action = strings.TrimSuffix(utils.Snake(controller), "_controller"), utils.Snake(action)

		c.Layout = "layouts/base.html"
		c.LayoutSections = make(map[string]string)

		c.Data["menu"] = ""
		c.Data["expand"] = ""

		c.LayoutSections["Styles"] = fmt.Sprintf("%s/%s_styles.%s", controller, action, "html")
		c.LayoutSections["Scripts"] = fmt.Sprintf("%s/%s_scripts.%s", controller, action, "html")
	}
}
