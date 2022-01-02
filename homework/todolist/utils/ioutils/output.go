/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 17:12:26
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 17:14:20
 */
package ioutils

import "fmt"

func Error(txt string) {
	fmt.Printf("[-] %s\n", txt)
}

func Success(txt string) {
	fmt.Printf("[+] %s\n", txt)
}

func Output(txt string) {
	fmt.Println(txt)
}
