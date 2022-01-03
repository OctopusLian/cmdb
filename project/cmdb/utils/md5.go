/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 21:49:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 21:49:10
 */
package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}
