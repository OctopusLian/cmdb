/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:08:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:12:18
 */
package cycle

import (
	"time"

	"agent/config"
	"agent/entity"
)

type Register struct {
	conf     *config.Config
	interval time.Duration
	nextTime time.Time
}

func (p *Register) Name() string {
	return "register"
}

func (p *Register) Init(conf *config.Config) {
	p.conf = conf
	p.interval = time.Hour
	// p.interval = time.Second * 30
	p.nextTime = time.Now()
}

func (p *Register) NextTime() time.Time {
	return p.nextTime
}

func (p *Register) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewRegister(p.conf.UUID), nil
}

func (p *Register) Pipline() chan interface{} {
	return p.conf.Register
}
