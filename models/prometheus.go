/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 23:25:48
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-13 12:57:01
 */
package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

type Node struct {
	ID        int
	UUID      string
	Hostname  string
	Addr      string
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Jobs      []*Job
}

type Jobs struct {
	ID        int
	Key       string
	Remark    string
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
}

type Target struct {
	ID        int
	UUID      string
	Hostname  string
	Addr      string
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Job       *Job
}

type Job struct {
}

func init() {
	orm.RegisterModel(new(Node), new(Jobs), new(Target))
}
