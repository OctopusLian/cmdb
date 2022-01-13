/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 18:44:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 18:51:47
 */
package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "mysql123"),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultString("mysql::Port", "3306"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

}
