/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 16:53:26
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:29:32
 */
package command

type Callback func()

type Command struct {
	Name     string
	Callback Callback
}

func New(name string, callback Callback) *Command {
	return &Command{
		Name:     name,
		Callback: callback,
	}
}
