/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 16:30:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 17:11:26
 */
package main

import (
	_ "cmdb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
