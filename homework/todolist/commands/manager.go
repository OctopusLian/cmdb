/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 16:52:43
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:38:09
 */
package commands

import (
	"fmt"
	"strconv"

	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/commands/command"
	"github.com/OctopusLian/Go-Operational-Development-Architecture/homework/todolist/utils/ioutils"
)

type manager struct {
	loginCallback command.LoginCallback
	cmds          map[int]*command.Command
}

func newManger() *manager {
	return &manager{
		cmds: make(map[int]*command.Command),
	}
}

func (mgr *manager) register(name string, callback command.Callback) {
	mgr.cmds[len(mgr.cmds)+1] = command.New(name, callback)
}

func (mgr *manager) registerLoginCallback(callback command.LoginCallback) {
	mgr.loginCallback = callback
}

func (mgr *manager) prompt() {
	for i := 1; i <= len(mgr.cmds); i++ {
		fmt.Printf("%d. %s", i, mgr.cmds[i].Name)
	}
}

func (mgr *manager) get(key int) (command.Callback, error) {
	if cmd, ok := mgr.cmds[key]; ok {
		return cmd.Callback, nil
	}
	return nil, fmt.Errorf("指令不存在")
}

func (mgr *manager) run() {
	for {
		mgr.prompt()
		key, err := strconv.Atoi(ioutils.Input("请输入指令： "))
		if err != nil {
			ioutils.Error("输入指令错误")
			continue
		}
		if callback, err := mgr.get(key); err != nil {
			ioutils.Error(err.Error())
		} else {
			callback()
		}
	}
}

var mgr *manager = newManger()

func RegisterLoginCallback(callback command.LoginCallback) {
	mgr.registerLoginCallback(callback)
}

func Register(name string, callback command.Callback) {
	mgr.register(name, callback)
}

func Run() {
	mgr.run()
}
