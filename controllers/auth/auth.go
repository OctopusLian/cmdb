package auth

import (
	"cmdb/controllers/base"
	"cmdb/models"
)

type LoginRequiredController struct {
	base.BaseController
	User *models.User
}

func (c *LoginRequiredController) Prepare() {
	c.BaseController.Prepare()
	if user := DefaultManager.IsLogin(c); user == nil {
		DefaultManager.GoLoginPage(c.Ctx, c.Ctx.Input.URL())
		c.StopRun()
	} else {
		c.User = user
		c.Data["user"] = user
	}
}

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	if DefaultManager.Login(c) {
		DefaultManager.GoHomePage(c.Ctx)
	}
}

func (c *AuthController) Logout() {
	DefaultManager.Logout(c)
	DefaultManager.GoLoginPage(c.Ctx, "")
}
