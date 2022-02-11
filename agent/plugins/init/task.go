/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:08:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:12:56
 */
package init

import (
	"agent/plugins"
	"agent/plugins/task"
)

func init() {
	plugins.DefaultManager.RegisterTask(&task.Process{})
}
