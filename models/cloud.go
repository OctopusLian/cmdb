package models

import (
	"time"

	_ "cmdb/cloud/plugins"

	"github.com/astaxie/beego/orm"
)

var PlatformTypes = map[string]string{
	"aws":       "AWS",
	"aliyun":    "阿里云",
	"tenantyun": "腾讯云",
}

type Platform struct {
	Model
	Id         int        `orm:"column(id);" json:"id"`
	Name       string     `orm:"column(name);size(64);" json:"name"`
	Type       string     `orm:"column(type);size(64);" json:"type"`
	Addr       string     `orm:"column(addr);size(1024);" json:"addr"`
	Region     string     `orm:"column(region);" json:"region"`
	Key        string     `orm:"column(key);size(1024);" json:"-"`
	Secrect    string     `orm:"column(secrect);size(1024);" json:"-"`
	Remark     string     `orm:"column(remark);size(1024);" json:"remark"`
	SyncTime   *time.Time `orm:"column(sync_time);type(datetime);null;" json:"sync_time"`
	Status     int        `orm:"column(status);default(0);" json:"status"`
	CreateUser int        `orm:"column(create_user);default(0);" json:"create_user"`

	Vms []*VirtualMachine `orm:"reverse(many)" json:"vms"`
}

func (u *Platform) IsEnable() bool {
	return u.Status == statusEnable
}

func (u *Platform) Enable(enable bool) bool {
	if enable {
		u.Status = statusEnable
	} else {
		u.Status = statusDisable
	}
	return true
}

func (u *Platform) AllEnabled() []*Platform {
	var list []*Platform
	ormer := orm.NewOrm()
	ormer.QueryTable(u).Filter("delete_time__isnull", true).Filter("status__exact", statusEnable).All(&list)
	return list
}

type VirtualMachine struct {
	Model
	Id            int       `orm:"column(id);" json:"id"`
	Key           string    `orm:"column(key);size(64);" json:"key"`
	UUID          string    `orm:"column(uuid);size(64);" json:"uuid"`
	Name          string    `orm:"column(name);size(128);" json:"name"`
	OS            string    `orm:"column(os);size(128);" json:"os"`
	CPU           int       `orm:"column(cpu);" json:"cpu"`
	Memory        int       `orm:"column(memory);" json:"memory"`
	PublicAddrs   string    `orm:"column(public_addrs);size(1024);" json:"public_addrs"`
	PrivateAddrs  string    `orm:"column(private_addrs);size(1024);" json:"private_addrs"`
	Status        string    `orm:"column(status);size(16);" json:"status"`
	VmCreatedTime string    `orm:"column(vm_created_time);" json:"vm_created_time"`
	VmExpiredTime string    `orm:"column(vm_expired_time);" json:"vm_expired_time"`
	Platform      *Platform `orm:"column(platform);rel(fk);" json:"platform"`
}

func init() {
	orm.RegisterModel(new(Platform), new(VirtualMachine))
}
