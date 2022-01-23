package forms

import (
	"strings"
	"time"

	"cmdb/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type UserCreateForm struct {
	Name           string    `form:"name"`
	Password       string    `form:"password"`
	PasswordVerify string    `form:"passwordVerify"`
	Gender         int       `form:"gender"`
	Birthday       time.Time `form:"birthday"`
	Department     string    `form:"department"`
	Tel            string    `form:"tel"`
	Email          string    `form:"email"`
	Addr           string    `form:"addr"`
	Remark         string    `form:"remark"`
}

func (f *UserCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)
	f.PasswordVerify = strings.TrimSpace(f.PasswordVerify)
	f.Department = strings.TrimSpace(f.Department)
	f.Tel = strings.TrimSpace(f.Tel)
	f.Email = strings.TrimSpace(f.Email)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := &models.User{Name: f.Name}
		if ormer.Read(user, "Name", "DeleteTime") != orm.ErrNoRows {
			v.SetError("name", "用户名已存在")
		}
	}

	v.MinSize(f.Password, 6, "password.password").Message("密码长度必须在%d-%d之内", 6, 32)
	v.MaxSize(f.Password, 32, "password.password").Message("密码长度必须在%d-%d之内", 6, 32)

	if f.PasswordVerify != f.PasswordVerify {
		v.SetError("passwordVerify", "两次输入密码不一致")
	}

	v.Range(f.Gender, 0, 1, "gender.gender").Message("性别选择不正确")

	v.MaxSize(f.Department, 512, "department.department").Message("部门长度必须在64个字符之内")

	v.Phone(f.Tel, "tel.tel").Message("电话格式不正确")
	v.Email(f.Email, "email.email").Message("邮箱格式不正确")

	v.MaxSize(f.Addr, 512, "addr.addr").Message("住址长度必须在512个字符之内")
	v.MaxSize(f.Remark, 512, "remark.remark").Message("备注长度必须在512个字符之内")
}

type UserModifyForm struct {
	Id         int       `form:"id"`
	Name       string    `form:"name"`
	Gender     int       `form:"gender"`
	Birthday   time.Time `form:"birthday"`
	Department string    `form:"department"`
	Tel        string    `form:"tel"`
	Email      string    `form:"email"`
	Addr       string    `form:"addr"`
	Remark     string    `form:"remark"`
}

func (f *UserModifyForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Department = strings.TrimSpace(f.Department)
	f.Tel = strings.TrimSpace(f.Tel)
	f.Email = strings.TrimSpace(f.Email)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)

	ormer := orm.NewOrm()
	user := &models.User{Id: f.Id}
	if ormer.Read(user) == orm.ErrNoRows {
		v.SetError("error", "操作对象不存在")
		return
	}

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := &models.User{Name: f.Name}
		if ormer.Read(user, "Name", "DeleteTime") != orm.ErrNoRows && user.Id != f.Id {
			v.SetError("name", "用户名已存在")
		}
	}

	v.Range(f.Gender, 0, 1, "gender.gender").Message("性别选择不正确")

	v.MaxSize(f.Department, 64, "department.department").Message("部门长度必须在64个字符之内")

	v.Phone(f.Tel, "tel.tel").Message("电话格式不正确")
	v.Email(f.Email, "email.email").Message("邮箱格式不正确")

	v.MaxSize(f.Addr, 512, "addr.addr").Message("住址长度必须在512个字符之内")
	v.MaxSize(f.Remark, 512, "remark.remark").Message("备注长度必须在512个字符之内")
}

type UserPasswordForm struct {
	OldPassword    string `form:"oldPassword"`
	Password       string `form:"password"`
	PasswordVerify string `form:"passwordVerify"`

	User *models.User
}

func (f *UserPasswordForm) Valid(v *validation.Validation) {
	f.OldPassword = strings.TrimSpace(f.OldPassword)
	f.Password = strings.TrimSpace(f.Password)
	f.PasswordVerify = strings.TrimSpace(f.PasswordVerify)

	if !f.User.ValidatePassword(f.OldPassword) {
		v.SetError("oldPassword", "密码不正确")
	}

	v.MinSize(f.Password, 6, "password.password").Message("密码长度必须在%d-%d之内", 6, 32)
	v.MaxSize(f.Password, 32, "password.password").Message("密码长度必须在%d-%d之内", 6, 32)

	if f.PasswordVerify != f.PasswordVerify {
		v.SetError("passwordVerify", "两次输入密码不一致")
	}
}
