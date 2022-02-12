/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-12 23:25:48
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-12 17:49:28
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Node struct {
	ID        int        `orm:"column(id);"`
	UUID      string     `orm:"column(uuid);varchar(64)"`
	Hostname  string     `orm:"varchar(64)"`
	Addr      string     `orm:"varchar(512)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Jobs      []*Job
}

type Job struct {
	ID        int        `orm:"column(id);"`
	Key       string     `orm:"varchar(64)"`
	Remark    string     `orm:"varchar(512)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Node      *Node      `orm:"rel(fk)"`
	Targets   []*Target  `orm:"reverse(many)"`
}

type Target struct {
	ID        int        `orm:"column(id);"`
	Name      string     `orm:"varchar(64)"`
	Remark    string     `orm:"varchar(512)"`
	Addr      string     `orm:"varchar(126)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Job       *Job       `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Node), new(Job), new(Target))
}
