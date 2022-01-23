package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"cmdb/models"
	"cmdb/utils"
)

var initForce bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "gocmdb init",
	Long:  `gocmdb init`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 设置日志到文件
		beego.SetLogger("file", `{
			"filename" : "logs/init.log",
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
		orm.RunSyncdb("default", initForce, verbose)
		ormer := orm.NewOrm()
		admin := &models.User{Name: "admin", IsSuperuser: true}
		if err := ormer.Read(admin, "Name"); err == orm.ErrNoRows {
			password := utils.RandString(6)
			admin.SetPassword(password)
			if _, err := ormer.Insert(admin); err == nil {
				beego.Informational("初始化admin成功, 默认密码:", password)
			} else {
				beego.Error("初始化用户失败, 错误:", err)
			}
		} else {
			beego.Informational("admin用户已存在, 跳过")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "force sync db(drop table)")
}
