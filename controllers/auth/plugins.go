package auth

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"cmdb/forms"
	"cmdb/models"
)

type Session struct {
}

func (s *Session) name() string {
	return "session"
}

func (s *Session) is(c *context.Context) bool {
	return c.Input.Header("Authentication") == ""
}

func (s *Session) login(c *AuthController) bool {
	form := &forms.LoginForm{Next: c.GetString("next")}
	valid := &validation.Validation{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
			} else if ok {
				c.SetSession("user", form.User.Id)
				return true
			}
		} else {
			valid.SetError("error", err.Error())
		}
	}
	c.TplName = "auth/login.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid
	return false
}

func (s *Session) isLogin(c *LoginRequiredController) *models.User {
	if session := c.GetSession("user"); session != nil {
		if uid, ok := session.(int); ok {
			user := models.User{Id: uid}
			ormer := orm.NewOrm()
			if ormer.Read(&user) == nil {
				return &user
			}
		}
	}
	return nil
}

func (s *Session) logout(c *AuthController) bool {
	c.DestroySession()
	return true
}

func (s *Session) goLoginPage(c *context.Context, next string) bool {
	if c.Input.IsAjax() {
		c.Output.JSON(map[string]interface{}{"code": 403, "text": "未认证"}, true, false)
	} else {
		c.Redirect(http.StatusFound, beego.URLFor(beego.AppConfig.DefaultString("login", "AuthController.Login"), "next", next))
	}
	return true
}

func (s *Session) goHomePage(c *context.Context) bool {
	next := c.Input.Query("next")
	if next == "" {
		next = beego.URLFor(beego.AppConfig.DefaultString("home", "HomeController.Index"))
	}
	c.Redirect(http.StatusFound, next)
	return true
}

type Token struct {
}

func (s *Token) name() string {
	return "token"
}

func (s *Token) is(c *context.Context) bool {
	return strings.ToLower(c.Input.Header("Authentication")) == "token"
}

func (t *Token) login(c *AuthController) bool {
	return true
}

func (t *Token) isLogin(c *LoginRequiredController) *models.User {
	c.EnableXSRF = false
	accessKey := c.Ctx.Input.Header("ACCESSKEY")
	signature := c.Ctx.Input.Header("SIGNATURE")
	ormer := orm.NewOrm()
	token := &models.Token{AccessKey: accessKey}
	if ormer.Read(token, "AccessKey") == nil && token.ValidateSignature(signature, c.Ctx.Input) {
		user := &models.User{Id: token.CreateUser}
		if ormer.Read(user) == nil && !user.IsLock() {
			return user
		}
	}
	return nil
}

func (t *Token) logout(c *AuthController) bool {
	return true
}

func (s *Token) goLoginPage(c *context.Context, next string) bool {
	c.Output.JSON(map[string]interface{}{"code": 400, "text": "请指定Token进行请求"}, true, false)
	return true
}

func (s *Token) goHomePage(c *context.Context) bool {
	c.Output.JSON(map[string]interface{}{"code": 400, "text": "请指定Token进行请求"}, true, false)
	return true
}

var defaultPlugin = new(Session)

func init() {
	DefaultManager.Register(new(Token))
}
