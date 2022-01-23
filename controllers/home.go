package controllers

import (
	"cmdb/controllers/auth"
	"net/http"

	"github.com/astaxie/beego"
)

type HomeController struct {
	auth.LoginRequiredController
}

func (c *HomeController) Index() {
	c.Redirect(beego.URLFor("UserPageController.Index"), http.StatusFound)
}
