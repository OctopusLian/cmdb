package controllers

import (
	"cmdb/controllers/auth"
	"cmdb/models"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare() {
	c.LoginRequiredController.Prepare()
	c.Layout = "layouts/base.html"

	c.LayoutSections = map[string]string{
		"LayoutStyle":  "",
		"LayoutScript": "",
	}

	c.Data["menu"] = ""
	c.Data["expand"] = ""

	alarmCount, alarms := models.DefaultAlarmManager.GetNotification(10)

	c.Data["alarm"] = map[string]interface{}{
		"count": alarmCount,
		"list":  alarms,
	}
}
