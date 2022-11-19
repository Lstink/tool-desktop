package device

import (
	"lstink.github.com/lstink/tool-desktop/internal/message"
)

type DeviceType int

const (
	XLDevice DeviceType = 1
	CXDevice DeviceType = 2
)

// ParseDataForMessage 解析提交过来的数据
func ParseDataForMessage(msg string, flag DeviceType) (res message.Data, errMsg string) {

	switch flag {
	case XLDevice:
		res, errMsg = xlParse(msg)
	case CXDevice:
		res, errMsg = cxParse(msg)
	}
	return
}

// 持续协议解析
func cxParse(msg string) (res message.Data, errMsg string) {

	return
}

// 小蓝协议解析
func xlParse(msg string) (res message.Data, errMsg string) {
	// 解析数据
	xl := message.NewXlMessage(msg)
	// 判断数据长度
	if !xl.Valid() {
		errMsg = "数据长度不正确"
		return
	}

	res = xl.GetParseData()

	return
}
