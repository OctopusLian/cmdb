/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-12 15:29:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 15:37:39
 */
package auth

import "strings"

//布局控制器基础
type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) getNav() string {
	controllerName, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}
