/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:17:09
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:18:17
 */
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"cmdb/cloud"
	_ "cmdb/cloud/plugins"
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
		for now := range time.Tick(10 * time.Second) {
			platforms, _, _ := models.DefaultCloudPlatformManager.Query("", 0, 0)
			for _, platform := range platforms {
				if !platform.IsEnable() {
					continue
				}
				if sdk, ok := cloud.DefaultManager.Cloud(platform.Type); !ok {
					beego.Error("云平台未注册")
				} else {
					sdk.Init(platform.Addr, platform.Region, platform.AccessKey, platform.SecrectKey)

					if err := sdk.TestConnect(); err != nil {
						beego.Error("测试链接失败:", err)
						models.DefaultCloudPlatformManager.SyncInfo(platform, now, fmt.Sprintf("测试链接失败: %s", err.Error()))
					} else {
						for _, instance := range sdk.GetInstance() {
							models.DefaultVirtualMachineManager.SyncInstance(instance, platform)
						}
						models.DefaultVirtualMachineManager.SyncInstanceStatus(now, platform)
						models.DefaultCloudPlatformManager.SyncInfo(platform, now, "")
					}
				}

			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloudCmd)
}
