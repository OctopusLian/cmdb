/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:08:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:13:10
 */
package plugins

import (
	"time"

	"agent/config"
)

type CyclePlugin interface {
	Name() string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipline() chan interface{}
}

type TaskPlugin interface {
	Name() string
	Init(*config.Config)
	Call(params string) (interface{}, error)
}
