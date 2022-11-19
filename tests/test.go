package main

import (
	"fmt"
	"lstink.github.com/lstink/tool-desktop/internal/message"
)

func main() {
	msg := "68150006004100000050000199020000000000000000028953"
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
