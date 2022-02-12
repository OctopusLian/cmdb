/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-12 18:01:15
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 18:09:52
 */
package cmd

import (
	"cmdb/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cobra"
)

var (
	name     string
	password string
)

var userCommand = &cobra.Command{
	Use:   "user",
	Short: "user console",
	Long:  "user console",
	RunE: func(cmd *cobra.Command, args []string) error {
		orm.Debug = verbose
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=True",
			beego.AppConfig.DefaultString("mysql::User", "root"),
			beego.AppConfig.DefaultString("mysql::Password", "mysql123"),
			beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
			beego.AppConfig.DefaultInt("mysql::Port", 3306),
		)

		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)

		//测试数据库连接
		if db, err := orm.GetDB("default"); err != nil {
			return err
		} else if err := db.Ping(); err != nil {
			return err
		}

		ormer := orm.NewOrm()
		user := &models.User{Name: name}
		user.Password = "123456"
		_, err := ormer.Insert(user)
		return err
	},
}

func init() {
	rootCmd.AddCommand(userCommand)
	userCommand.Flags().StringVarP(&name, "name", "n", "admin", "name")
	userCommand.Flags().StringVarP(&password, "password", "p", "123456", "password")
}
