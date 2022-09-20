package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) //返回命令行参数
	fmt.Println("11")
	if len(os.Args) > 1 {

	}
	os.Exit(144) //返回值
}
