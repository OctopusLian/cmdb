package models

import "time"

type Model struct {
	CreateTime *time.Time `orm:"column(create_time);type(datetime);auto_now_add;"  json:"create_time"`
	UpdateTime *time.Time `orm:"column(update_time);type(datetime);auto_now;" json:"update_time"`
	DeleteTime *time.Time `orm:"column(delete_time);type(datetime);null;"  json:"delete_time"`
}

func (m *Model) Patch() {

}

func (m *Model) Delete() bool {
	now := time.Now()
	m.DeleteTime = &now
	return true
}
