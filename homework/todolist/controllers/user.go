/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 17:24:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:47:54
 */
package controllers

import (
	"fmt"
	"os"

	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/config"
	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/utils/ioutils"
)

func Logout() {
	os.Exit(0)
}

func Login() bool {
	for i := config.Config.LoginRetry; i > 0; i-- {
		txt := ioutils.Input("请输入密码：")
		if txt == "kk" {
			return true
		}
		if i != 1 {
			ioutils.Error(fmt.Sprintf("密码错误", "剩余登录%d次数", i-1))
		}
	}

	ioutils.Error(fmt.Sprintf("密码错误超过%d, 程序退出", config.Config.LoginRetry))
	return false
}
