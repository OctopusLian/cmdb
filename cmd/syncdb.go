package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var syncDbForce bool

var syncDbCmd = &cobra.Command{
	Use:   "syncdb",
	Short: "gocmdb syncdb",
	Long:  `gocmdb syncdb`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 设置日志到文件
		beego.SetLogger("file", `{
			"filename" : "logs/syncdb.log",
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
		orm.RunSyncdb("default", syncDbForce, verbose)
		beego.Informational("同步数据库")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncDbCmd)
	syncDbCmd.Flags().BoolVarP(&syncDbForce, "force", "f", false, "force sync db(drop table)")
}
