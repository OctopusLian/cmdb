/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 17:32:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:32:45
 */
package init

import (
	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/commands"
	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/controllers"
)

func init() {
	commands.Register("退出", controllers.Logout)
}
