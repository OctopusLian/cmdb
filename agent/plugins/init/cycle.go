/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:08:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:12:38
 */
package init

import (
	"agent/plugins"
	"agent/plugins/cycle"
)

func init() {
	plugins.DefaultManager.RegisterCycle(&cycle.Heartbeat{})
	plugins.DefaultManager.RegisterCycle(&cycle.Register{})
	plugins.DefaultManager.RegisterCycle(&cycle.Resource{})
}
