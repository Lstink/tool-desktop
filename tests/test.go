package main

import (
	"fmt"
	"lstink.github.com/lstink/tool-desktop/internal/message"
)

func main() {
	msg := "685e0200000a00000050000242010060ea00003075000060ea00003075000060ea00003075000060ea0000307500000003030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030333e5"
	// 解析数据
	xl := message.NewXlMessage(msg)
	// 判断数据长度
	if (xl.Length+4)*2 != len(msg) {
		fmt.Println("数据长度不正确")
	}

	xl.GetParseData()

}
