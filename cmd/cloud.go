package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"cmdb/cloud"
	"cmdb/models"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "gocmdb cloud",
	Long:  `gocmdb cloud`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 设置日志到文件
		beego.SetLogger("file", `{
			"filename" : "logs/cloud.log",
			"level" : 7}`,
		)
		if !verbose {
			//删除控制台日志
			beego.BeeLogger.DelLogger("console")
		} else {
			orm.Debug = true
		}

		// 初始化orm
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

		// 测试数据库连接是否正常
		if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
			beego.Error("数据库连接错误")
			return fmt.Errorf("数据库连接错误")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		platform := new(models.Platform)
		ormer := orm.NewOrm()
		for now := range time.Tick(10 * time.Second) {
			for _, pf := range platform.AllEnabled() {
				go func(pf *models.Platform) {
					beego.Debug(pf.Name)
					sdk, err := cloud.DefaultManager.Cloud(pf.Type)
					if err != nil {
						beego.Error(err)
						return
					}

					sdk.Init(pf.Addr, pf.Key, pf.Secrect, pf.Region)
					if err := sdk.TestConnect(); err != nil {
						beego.Error(err)
						return
					}

					for _, instance := range sdk.GetInstances() {
						obj := &models.VirtualMachine{UUID: instance.UUID, Platform: pf}
						if _, _, err := ormer.ReadOrCreate(obj, "UUID", "Platform"); err != nil {
							beego.Error(err)
							continue
						}
						obj.Key = instance.Key
						obj.Name = instance.Name
						obj.OS = instance.OS
						obj.CPU = instance.CPU
						obj.Memory = instance.Memory
						obj.PublicAddrs = strings.Join(instance.PublicAddrs, ",")
						obj.PrivateAddrs = strings.Join(instance.PrivateAddrs, ",")
						obj.Status = instance.Status
						obj.VmCreatedTime = instance.CreatedTime
						obj.VmExpiredTime = instance.ExpiredTime
						ormer.Update(obj)
					}
					ormer.QueryTable(new(models.VirtualMachine)).Filter("platform__exact", pf).Filter("update_time__gte", now).Update(orm.Params{"delete_time": nil})
					ormer.QueryTable(new(models.VirtualMachine)).Filter("platform__exact", pf).Filter("update_time__lt", now).Update(orm.Params{"delete_time": now})
					pf.SyncTime = &now
					ormer.Update(pf, "sync_time")
				}(pf)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloudCmd)
}
