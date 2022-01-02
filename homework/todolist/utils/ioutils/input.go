/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 17:07:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:51:23
 */
package ioutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func Password(prompt string) string {
	fmt.Print(prompt)
	if ctx, err := gopass.GetPasswd(); err != nil {
		return ""
	} else {
		return strings.TrimSpace(string(ctx))
	}
}
