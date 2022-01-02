/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 12:35:21
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 16:46:04
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func tailf(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	//seek文件末尾
	file.Seek(0, os.SEEK_END)
	reader := bufio.NewReader(file)
	//var builder strings.Builder
	for {
		line, _, _ := reader.ReadLine()
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(string(line))
		//文件无换行，先放入缓冲
		// if line[len(line)-1] == '\n' {
		// 	builder.String()
		//fmt.Fprintf(os.Stdout,"%s%s",builder.String(),line)
		// 	fmt.Println(line)
		// 	builder.Reset() //清空
		// } else {
		// 	builder.Write(line)
		// }
	}
}

func main() {
	var path string
	var h, help bool
	flag.StringVar(&path, "p", "tail.log", "path") //读取本地默认文件
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")

	flag.Usage = func() {
		fmt.Println("tailf -p path")
		flag.PrintDefaults()
	}

	flag.Parse()
	if h || help {
		flag.Usage()
		os.Exit(0)
	}

	//检查path
	tailf(path)
}
