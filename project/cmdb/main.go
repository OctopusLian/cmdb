/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 16:30:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 19:04:05
 */
package main

import (
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"

	_ "cmdb/routers"
)

func main() {
	beego.Run()
}
