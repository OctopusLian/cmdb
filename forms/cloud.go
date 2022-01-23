package forms

import (
	"strings"

	"cmdb/cloud"
	"cmdb/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type PlatformCreateForm struct {
	Name    string `form:"name"`
	Type    string `form:"type"`
	Addr    string `form:"addr"`
	Region  string `form:"region"`
	Key     string `form:"key"`
	Secrect string `form:"secrect"`
	Remark  string `form:"remark"`
}

func (f *PlatformCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Type = strings.TrimSpace(f.Type)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Region = strings.TrimSpace(f.Region)
	f.Key = strings.TrimSpace(f.Key)
	f.Secrect = strings.TrimSpace(f.Secrect)
	f.Remark = strings.TrimSpace(f.Remark)

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		platform := &models.Platform{Name: f.Name}
		if ormer.Read(platform, "Name", "DeleteTime") != orm.ErrNoRows {
			v.SetError("name", "名称已存在")
		}
	}

	if _, ok := models.PlatformTypes[f.Type]; !ok {
		v.SetError("type", "平台选择不正确")
	}

	v.MaxSize(f.Addr, 512, "addr.addr").Message("地址长度必须在512个字符之内")

	v.MinSize(f.Region, 1, "region.region").Message("区域长度必须在%d-%d之内", 1, 32)
	v.MaxSize(f.Region, 32, "region.region").Message("区域长度必须在%d-%d之内", 1, 32)

	v.MinSize(f.Key, 1, "key.key").Message("Key长度必须在%d-%d之内", 1, 512)
	v.MaxSize(f.Key, 512, "key.key").Message("Key长度必须在%d-%d之内", 1, 512)

	v.MinSize(f.Secrect, 1, "secrect.secrect").Message("Secrect长度必须在%d-%d之内", 1, 512)
	v.MaxSize(f.Secrect, 512, "secrect.secrect").Message("Secrect长度必须在%d-%d之内", 1, 512)

	v.MaxSize(f.Remark, 512, "remark.remark").Message("备注长度必须在512个字符之内")

	if !v.HasErrors() {
		plugin, err := cloud.DefaultManager.Cloud(f.Type)
		if err != nil {
			v.SetError("error", err.Error())
		} else {
			plugin.Init(f.Addr, f.Key, f.Secrect, f.Region)
			if err := plugin.TestConnect(); err != nil {
				v.SetError("error", "配置不正确, 测试连接失败")
			}
		}
	}
}

type PlatformModifyForm struct {
	Id      int    `form:"id"`
	Name    string `form:"name"`
	Type    string `form:"type"`
	Addr    string `form:"addr"`
	Region  string `form:"region"`
	Key     string `form:"key"`
	Secrect string `form:"secrect"`
	Remark  string `form:"remark"`
}

func (f *PlatformModifyForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Type = strings.TrimSpace(f.Type)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Region = strings.TrimSpace(f.Region)
	f.Key = strings.TrimSpace(f.Key)
	f.Secrect = strings.TrimSpace(f.Secrect)
	f.Remark = strings.TrimSpace(f.Remark)

	ormer := orm.NewOrm()
	platform := &models.Platform{Id: f.Id}
	if ormer.Read(platform) == orm.ErrNoRows {
		v.SetError("error", "操作对象不存在")
		return
	}

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		platform := &models.Platform{Name: f.Name}
		if ormer.Read(platform, "Name", "DeleteTime") != orm.ErrNoRows && platform.Id != f.Id {
			v.SetError("name", "名称已存在")
		}
	}

	if _, ok := models.PlatformTypes[f.Type]; !ok {
		v.SetError("type", "平台选择不正确")
	}

	v.MaxSize(f.Addr, 512, "addr.addr").Message("地址长度必须在512个字符之内")

	v.MinSize(f.Region, 1, "region.region").Message("区域长度必须在%d-%d之内", 1, 32)
	v.MaxSize(f.Region, 32, "region.region").Message("区域长度必须在%d-%d之内", 1, 32)

	v.MinSize(f.Key, 0, "key.key").Message("Key长度必须在%d-%d之内", 0, 512)
	v.MaxSize(f.Key, 512, "key.key").Message("Key长度必须在%d-%d之内", 0, 512)

	v.MinSize(f.Secrect, 0, "secrect.secrect").Message("Secrect长度必须在%d-%d之内", 0, 512)
	v.MaxSize(f.Secrect, 512, "secrect.secrect").Message("Secrect长度必须在%d-%d之内", 0, 512)

	v.MaxSize(f.Remark, 512, "remark.remark").Message("备注长度必须在512个字符之内")

	if !v.HasErrors() {
		plugin, err := cloud.DefaultManager.Cloud(f.Type)
		if err != nil {
			v.SetError("error", err.Error())
		} else {
			key := f.Key
			if key == "" {
				key = platform.Key
			}
			secrect := f.Secrect
			if secrect == "" {
				secrect = platform.Secrect
			}
			plugin.Init(f.Addr, key, secrect, f.Region)
			if err := plugin.TestConnect(); err != nil {
				v.SetError("error", "配置不正确, 测试连接失败")
			}
		}
	}
}
