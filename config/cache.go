/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-15 23:30:00
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 23:32:39
 */
package config

import (
	"fmt"

	"github.com/beego/beego/v2/client/cache"
)

var Cache cache.Cache

func Init(adapter, config string) {
	var err error
	Cache, err := cache.NewCache(adapter, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(Cache) //TODO
}
