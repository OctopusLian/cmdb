/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 17:07:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:27:48
 */
package ioutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
