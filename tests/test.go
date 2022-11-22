package main

import (
	"fmt"
	"lstink.github.com/lstink/tool-desktop/internal/message"
)

func main() {
	msg := "685e00250058000000123456780100a086010000000000a086010000000000a086010000000000a08601000000000000030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303ba23"
	// 解析数据
	xl := message.NewXlMessage(msg)
	// 判断数据长度
	if (xl.Length+4)*2 != len(msg) {
		fmt.Println("数据长度不正确")
	}
	// 解析数据
	res := xl.GetParseData()
	fmt.Println(res)

}
