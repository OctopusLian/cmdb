package forms

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"cmdb/models"
)

type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Next     string `form:"next"`

	User *models.User
}

func (f *LoginForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)

	ormer := orm.NewOrm()
	user := &models.User{Name: f.Name}
	if ormer.Read(user, "Name") != nil || !user.ValidatePassword(f.Password) {
		v.SetError("login", "用户名或密码错误")
	} else if user.IsLock() {
		v.SetError("login", "用户已被锁定")
	} else {
		f.User = user
	}
}
