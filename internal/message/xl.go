package message

import (
	"fmt"
	"lstink.github.com/lstink/tool-desktop/utils"
	"strings"
)

const (
	XlRemoteConfirmBicycle = 0x72
	XlOrderLogBicycle      = 0x79
	XlGetBalanceForBusBack = 0x0a
)

type HexParse int

const (
	Bin    HexParse = 1
	Bcd    HexParse = 2
	CpTIme HexParse = 3
	Card1  HexParse = 4
	Card2  HexParse = 5
)

type XlMessage struct {
	Head      string
	Length    int
	Inc       string
	Flag      string
	Command   int
	Data      string
	Crc       string
	OriginMsg string
}

// NewXlMessage 实例化
func NewXlMessage(data string) *XlMessage {
	return &XlMessage{
		Head:      data[:2],
		Length:    utils.HexToTen(data[2:4]),
		Inc:       data[4:8],
		Flag:      data[8:10],
		Command:   utils.HexToTen(data[10:12]),
		Data:      data[12 : len(data)-4],
		Crc:       data[len(data)-4:],
		OriginMsg: data,
	}
}

// Valid 校验数据
func (xl XlMessage) Valid() bool {
	// 判断数据长度
	if (xl.Length+4)*2 != len(xl.OriginMsg) {
		return false
	}

	return true
}

// GetParseData 获取解析到的数据
func (xl XlMessage) GetParseData() (data []Data) {
	switch xl.Command {
	case XlOrderLogBicycle:
		fmt.Println("上传订单")
		data = xl.OrderLogBicycle()
	case XlGetBalanceForBusBack:
		fmt.Println("充电桩计费模型请求(汽车桩)回复")
		data = xl.GetBalanceForBusBack()
	}

	return
}

func (xl XlMessage) OrderLogBicycle() (data []Data) {
	// 解析数据
	data = append(data, Data{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data = append(data, Data{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data = append(data, Data{Key: "枪号", Value: strings.TrimLeft(xl.Data[46:48], "0")})
	data = append(data, Data{Key: "开始时间", Value: getParseData(xl.Data[48:62], CpTIme)})
	data = append(data, Data{Key: "结束时间", Value: getParseData(xl.Data[62:76], CpTIme)})
	data = append(data, Data{Key: "账户类型", Value: getParseData(xl.Data[76:78], Bin)})
	data = append(data, Data{Key: "计费方式", Value: getParseData(xl.Data[78:80], Bin)})
	data = append(data, Data{Key: "启动模式", Value: getParseData(xl.Data[80:82], Bin)})
	data = append(data, Data{Key: "启动模式参数", Value: getParseData(xl.Data[82:84], Bin)})
	data = append(data, Data{Key: "逻辑卡号", Value: getParseData(xl.Data[86:102], Card1)})
	data = append(data, Data{Key: "结束时功率", Value: getParseData(xl.Data[102:106], Bin)})
	data = append(data, Data{Key: "电表总起值", Value: getParseData(xl.Data[106:114], Bin)})
	data = append(data, Data{Key: "电表总止值", Value: getParseData(xl.Data[114:122], Bin)})
	data = append(data, Data{Key: "总电量", Value: getParseData(xl.Data[122:130], Bin)})
	data = append(data, Data{Key: "计损总电量", Value: getParseData(xl.Data[130:138], Bin)})
	data = append(data, Data{Key: "消费金额", Value: getParseData(xl.Data[138:146], Bin)})
	data = append(data, Data{Key: "SOC", Value: getParseData(xl.Data[146:148], Bin)})
	data = append(data, Data{Key: "交易标识", Value: getParseData(xl.Data[148:150], Bin)})
	data = append(data, Data{Key: "交易时间", Value: getParseData(xl.Data[150:164], CpTIme)})
	data = append(data, Data{Key: "停止原因", Value: getParseData(xl.Data[164:166], Bin)})

	return
}

func (xl XlMessage) GetBalanceForBusBack() (data []Data) {
	// 解析数据
	data = append(data, Data{Key: "桩编号", Value: strings.TrimLeft(xl.Data[0:14], "0")})
	data = append(data, Data{Key: "计费模型编号", Value: getParseData(xl.Data[14:18], Bin)})
	data = append(data, Data{Key: "尖费电费费率", Value: getParseData(xl.Data[18:26], Bin)})
	data = append(data, Data{Key: "尖服务费费率", Value: getParseData(xl.Data[26:34], Bin)})
	data = append(data, Data{Key: "峰电费费率", Value: getParseData(xl.Data[34:42], Bin)})
	data = append(data, Data{Key: "峰服务费费率", Value: getParseData(xl.Data[42:50], Bin)})
	data = append(data, Data{Key: "平电费费率", Value: getParseData(xl.Data[50:58], Bin)})
	data = append(data, Data{Key: "平服务费费率", Value: getParseData(xl.Data[58:66], Bin)})
	data = append(data, Data{Key: "谷电费费率", Value: getParseData(xl.Data[66:74], Bin)})
	data = append(data, Data{Key: "谷服务费费率", Value: getParseData(xl.Data[74:82], Bin)})
	data = append(data, Data{Key: "计损比例", Value: getParseData(xl.Data[82:84], Bin)})
	data = append(data, Data{Key: "0：00～0：30时段费率号", Value: getParseData(xl.Data[84:86], Bin)})
	data = append(data, Data{Key: "0：30～1：00时段费率号", Value: getParseData(xl.Data[100:102], Bin)})
	data = append(data, Data{Key: "......", Value: getParseData(xl.Data[102:104], Bin)})
	data = append(data, Data{Key: "23：00～23：30时段费率号", Value: getParseData(xl.Data[104:106], Bin)})
	data = append(data, Data{Key: "23：30～0：00时段费率号", Value: getParseData(xl.Data[106:108], Bin)})

	return
}

// 解析数据
func getParseData(str string, f HexParse) (res string) {
	switch f {
	case Bin:
		res = utils.GetDataForBinCode(str, 10)
	case Bcd:
		res = str
	case CpTIme:
		res = utils.ChangeCp56Time2a(str)
	case Card1:
		res = utils.RevertTwoByte(str)
	case Card2:
		res = utils.RevertTwoByte(str)
	}

	return
}
