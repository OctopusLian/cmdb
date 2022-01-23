/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-22 23:58:29
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-23 19:10:53
 */
package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func Md5(txt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(txt)))
}

func Md5Salt(txt, salt string) string {
	if salt == "" {
		salt = RandString(8)
	}
	return fmt.Sprintf("%s:%s", salt, Md5(fmt.Sprintf("%s:%s", salt, txt)))
}

func SplitMd5Salt(txt string) (string, string) {
	elements := strings.SplitN(txt, ":", 2)
	if len(elements) > 1 {
		return elements[0], elements[1]
	}
	return elements[0], ""
}

// 计算字符串sha256 hex值
func Sha256Hex(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
}

// 计算字符串hmac-sha256值
func HS256(text, key string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(text))
	return string(hasher.Sum(nil))
}

// 计算字符串hmac-sha256 hex值
func HS256Hex(text, key string) string {
	return strings.ToLower(hex.EncodeToString([]byte(HS256(text, key))))
}
