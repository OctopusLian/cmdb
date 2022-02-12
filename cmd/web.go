/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:17:09
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 18:09:18
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "cmdb/cloud/plugins"
	"cmdb/config"
	_ "cmdb/routers"

	_ "github.com/astaxie/beego/session/redis"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Web console",
	Long:  "Web console",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 设置日志到文件
		beego.SetLogger("file", `{
			"filename" : "logs/cmdb.log",
			//// "level" : 7}`,
		)
		beego.SetLogFuncCall(true)
		beego.SetLevel(beego.LevelDebug)
		if !verbose {
			//删除控制台日志
			beego.BeeLogger.DelLogger("console")
		} else {
			orm.Debug = true
		}
		config.Init("file", `{"CachePath":"tmp/cache","FileSuffix":".cache",}`)
		orm.Debug = verbose
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=True",
			beego.AppConfig.DefaultString("mysql::User", "root"),
			beego.AppConfig.DefaultString("mysql::Password", "mysql123"),
			beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
			beego.AppConfig.DefaultInt("mysql::Port", 3306),
		)

		// 初始化orm
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)

		// 测试数据库连接是否正常
		if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
			beego.Error("数据库连接错误")
			return fmt.Errorf("数据库连接错误")
		}
		return nil
	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	beego.Run()
	// 	return nil
	// },
}

func init() {
	rootCmd.AddCommand(webCmd)
}
