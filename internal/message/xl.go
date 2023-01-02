package message

import (
	"fmt"
	"lstink.github.com/lstink/tool-desktop/utils"
	"strings"
)

const (
	//XlRemoteConfirmBicycle = 0x72
	XlOrderLogBicycle      = 0x79
	XlGetBalanceForBusBack = 0x0a
	XlOrderLogBus          = 0x39
	XlStartCharging        = 0x34
	XlUserMoneyUpdateBack  = 0x41
	XlUpdateUserMoney      = 0x42
	XlStartChargingBicycle = 0x74
	XlBalanceModelSetting  = 0x58
	XlBusRealTimeData      = 0x11
)

type HexParse int

const (
	Bin    HexParse = 1
	Bcd    HexParse = 2
	CpTIme HexParse = 3
	Card1  HexParse = 4
	Card2  HexParse = 5
	Ascii  HexParse = 6
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
func (xl XlMessage) GetParseData() (data Data) {
	switch xl.Command {
	case XlOrderLogBicycle:
		fmt.Println("上传单车订单")
		data = xl.OrderLogBicycle()
	case XlOrderLogBus:
		fmt.Println("上传汽车订单")
		data = xl.OrderLogBus()
	case XlGetBalanceForBusBack:
		fmt.Println("充电桩计费模型请求(汽车桩)回复")
		data = xl.GetBalanceForBusBack()
	case XlStartCharging:
		fmt.Println("运营平台远程启动汽车桩")
		data = xl.RemoteStartCharging()
	case XlUpdateUserMoney:
		fmt.Println("运营平台远程启动汽车桩")
		data = xl.XlUpdateUserMoney()
	case XlUserMoneyUpdateBack:
		fmt.Println("用户余额更新应答")
		data = xl.XlUserMoneyUpdateBack()
	case XlStartChargingBicycle:
		fmt.Println("小蓝单车启动充电")
		data = xl.XlStartChargingBicycle()
	case XlBalanceModelSetting:
		fmt.Println("交流桩计费模型设置")
		data = xl.XlBalanceModelSetting()
	case XlBusRealTimeData:
		fmt.Println("汽车桩实时数据")
		data = xl.XlBusRealTimeData()
	}

	return
}

func (xl XlMessage) OrderLogBicycle() (data Data) {

	data.Cmd = xl.Command
	data.Remark = "单车桩订单记录"
	// 解析数据
	data.List = append(data.List, Item{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: strings.TrimLeft(xl.Data[46:48], "0")})
	data.List = append(data.List, Item{Key: "开始时间", Value: getParseData(xl.Data[48:62], CpTIme)})
	data.List = append(data.List, Item{Key: "结束时间", Value: getParseData(xl.Data[62:76], CpTIme)})
	data.List = append(data.List, Item{Key: "账户类型", Value: getParseData(xl.Data[76:78], Bin)})
	data.List = append(data.List, Item{Key: "计费方式", Value: getParseData(xl.Data[78:80], Bin)})
	data.List = append(data.List, Item{Key: "启动模式", Value: getParseData(xl.Data[80:82], Bin)})
	data.List = append(data.List, Item{Key: "启动模式参数", Value: getParseData(xl.Data[82:84], Bin)})
	data.List = append(data.List, Item{Key: "逻辑卡号", Value: getParseData(xl.Data[86:102], Card1)})
	data.List = append(data.List, Item{Key: "结束时功率", Value: getParseData(xl.Data[102:106], Bin)})
	data.List = append(data.List, Item{Key: "电表总起值", Value: getParseData(xl.Data[106:114], Bin)})
	data.List = append(data.List, Item{Key: "电表总止值", Value: getParseData(xl.Data[114:122], Bin)})
	data.List = append(data.List, Item{Key: "总电量", Value: getParseData(xl.Data[122:130], Bin)})
	data.List = append(data.List, Item{Key: "计损总电量", Value: getParseData(xl.Data[130:138], Bin)})
	data.List = append(data.List, Item{Key: "消费金额", Value: getParseData(xl.Data[138:146], Bin)})
	data.List = append(data.List, Item{Key: "SOC", Value: getParseData(xl.Data[146:148], Bin)})
	data.List = append(data.List, Item{Key: "交易标识", Value: getParseData(xl.Data[148:150], Bin)})
	data.List = append(data.List, Item{Key: "交易时间", Value: getParseData(xl.Data[150:164], CpTIme)})
	data.List = append(data.List, Item{Key: "停止原因", Value: getParseData(xl.Data[164:166], Bin)})

	return
}

func (xl XlMessage) GetBalanceForBusBack() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "汽车桩下发计费模型"
	// 解析数据
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[0:14], "0")})
	data.List = append(data.List, Item{Key: "计费模型编号", Value: getParseData(xl.Data[14:18], Bin)})
	data.List = append(data.List, Item{Key: "尖费电费费率", Value: getParseData(xl.Data[18:26], Bin)})
	data.List = append(data.List, Item{Key: "尖服务费费率", Value: getParseData(xl.Data[26:34], Bin)})
	data.List = append(data.List, Item{Key: "峰电费费率", Value: getParseData(xl.Data[34:42], Bin)})
	data.List = append(data.List, Item{Key: "峰服务费费率", Value: getParseData(xl.Data[42:50], Bin)})
	data.List = append(data.List, Item{Key: "平电费费率", Value: getParseData(xl.Data[50:58], Bin)})
	data.List = append(data.List, Item{Key: "平服务费费率", Value: getParseData(xl.Data[58:66], Bin)})
	data.List = append(data.List, Item{Key: "谷电费费率", Value: getParseData(xl.Data[66:74], Bin)})
	data.List = append(data.List, Item{Key: "谷服务费费率", Value: getParseData(xl.Data[74:82], Bin)})
	data.List = append(data.List, Item{Key: "计损比例", Value: getParseData(xl.Data[82:84], Bin)})
	data.List = append(data.List, Item{Key: "0：00～0：30时段费率号", Value: getParseData(xl.Data[84:86], Bin)})
	data.List = append(data.List, Item{Key: "0：30～1：00时段费率号", Value: getParseData(xl.Data[100:102], Bin)})
	data.List = append(data.List, Item{Key: "......", Value: getParseData(xl.Data[102:104], Bin)})
	data.List = append(data.List, Item{Key: "23：00～23：30时段费率号", Value: getParseData(xl.Data[104:106], Bin)})
	data.List = append(data.List, Item{Key: "23：30～0：00时段费率号", Value: getParseData(xl.Data[106:108], Bin)})

	return
}

// RemoteStartCharging 汽车桩远程控制启机
func (xl XlMessage) RemoteStartCharging() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "汽车桩远程控制启机"
	// 解析数据
	data.List = append(data.List, Item{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: strings.TrimLeft(xl.Data[46:48], "0")})
	data.List = append(data.List, Item{Key: "逻辑卡号", Value: strings.TrimLeft(xl.Data[48:64], "0")})
	data.List = append(data.List, Item{Key: "物理卡号", Value: strings.TrimLeft(xl.Data[64:80], "0")})
	data.List = append(data.List, Item{Key: "账户余额", Value: getParseData(xl.Data[80:88], Bin)})

	return
}

func (xl XlMessage) OrderLogBus() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "汽车桩订单上传"
	// 解析数据
	data.List = append(data.List, Item{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: getParseData(xl.Data[46:48], Bin)})
	data.List = append(data.List, Item{Key: "开始时间", Value: getParseData(xl.Data[48:62], CpTIme)})
	data.List = append(data.List, Item{Key: "结束时间", Value: getParseData(xl.Data[62:76], CpTIme)})
	data.List = append(data.List, Item{Key: "尖单价", Value: getParseData(xl.Data[76:84], Bin)})
	data.List = append(data.List, Item{Key: "尖电量", Value: getParseData(xl.Data[84:92], Bin)})
	data.List = append(data.List, Item{Key: "计损尖电量", Value: getParseData(xl.Data[92:100], Bin)})
	data.List = append(data.List, Item{Key: "尖金额", Value: getParseData(xl.Data[100:108], Bin)})
	data.List = append(data.List, Item{Key: "峰单价", Value: getParseData(xl.Data[108:116], Bin)})
	data.List = append(data.List, Item{Key: "峰电量", Value: getParseData(xl.Data[116:124], Bin)})
	data.List = append(data.List, Item{Key: "计损峰电量", Value: getParseData(xl.Data[124:132], Bin)})
	data.List = append(data.List, Item{Key: "峰金额", Value: getParseData(xl.Data[132:140], Bin)})
	data.List = append(data.List, Item{Key: "平单价", Value: getParseData(xl.Data[140:148], Bin)})
	data.List = append(data.List, Item{Key: "平电量", Value: getParseData(xl.Data[148:156], Bin)})
	data.List = append(data.List, Item{Key: "计损平电量", Value: getParseData(xl.Data[156:164], Bin)})
	data.List = append(data.List, Item{Key: "平金额", Value: getParseData(xl.Data[164:172], Bin)})
	data.List = append(data.List, Item{Key: "谷单价", Value: getParseData(xl.Data[172:180], Bin)})
	data.List = append(data.List, Item{Key: "谷电量", Value: getParseData(xl.Data[180:188], Bin)})
	data.List = append(data.List, Item{Key: "计损谷电量", Value: getParseData(xl.Data[188:196], Bin)})
	data.List = append(data.List, Item{Key: "谷金额", Value: getParseData(xl.Data[196:204], Bin)})
	data.List = append(data.List, Item{Key: "电表总起值", Value: getParseData(xl.Data[204:212], Bin)})
	data.List = append(data.List, Item{Key: "电表总止值", Value: getParseData(xl.Data[212:220], Bin)})
	data.List = append(data.List, Item{Key: "总电量", Value: getParseData(xl.Data[220:228], Bin)})
	data.List = append(data.List, Item{Key: "计损总电量", Value: getParseData(xl.Data[228:236], Bin)})
	data.List = append(data.List, Item{Key: "消费金额", Value: getParseData(xl.Data[236:244], Bin)})
	data.List = append(data.List, Item{Key: "电动汽车唯一标识", Value: getParseData(xl.Data[244:278], Ascii)})
	data.List = append(data.List, Item{Key: "交易标识", Value: getParseData(xl.Data[278:280], Bin)})
	data.List = append(data.List, Item{Key: "交易日期、时间", Value: getParseData(xl.Data[280:294], CpTIme)})
	data.List = append(data.List, Item{Key: "停止原因", Value: getParseData(xl.Data[294:296], Bin)})
	data.List = append(data.List, Item{Key: "物理卡号", Value: getParseData(xl.Data[296:312], Bin)})

	return
}

// XlUpdateUserMoney 更新用户余额
func (xl XlMessage) XlUpdateUserMoney() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "更新用户余额"
	// 解析数据
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[0:14], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: strings.TrimLeft(xl.Data[14:16], "0")})
	data.List = append(data.List, Item{Key: "物理卡号", Value: strings.TrimLeft(xl.Data[16:32], "0")})
	data.List = append(data.List, Item{Key: "账户余额", Value: getParseData(xl.Data[32:40], Bin)})

	return
}

// XlUserMoneyUpdateBack 更新用户余额应答
func (xl XlMessage) XlUserMoneyUpdateBack() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "更新用户余额应答"
	// 解析数据
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[0:14], "0")})
	data.List = append(data.List, Item{Key: "物理卡号", Value: strings.TrimLeft(xl.Data[14:30], "0")})
	data.List = append(data.List, Item{Key: "修改结果", Value: getParseData(xl.Data[30:32], Bin)})

	return
}

// XlStartChargingBicycle 小蓝单车远程控制启机
func (xl XlMessage) XlStartChargingBicycle() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "小蓝单车远程控制启机"
	// 解析数据
	data.List = append(data.List, Item{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: strings.TrimLeft(xl.Data[46:48], "0")})
	data.List = append(data.List, Item{Key: "逻辑卡号", Value: strings.TrimLeft(xl.Data[48:64], "0")})
	data.List = append(data.List, Item{Key: "账户余额", Value: getParseData(xl.Data[64:72], Bin)})
	data.List = append(data.List, Item{Key: "充电模式参数", Value: getParseData(xl.Data[72:76], Bin)})
	data.List = append(data.List, Item{Key: "账户类型", Value: getParseData(xl.Data[76:78], Bin)})
	data.List = append(data.List, Item{Key: "计费方式", Value: getParseData(xl.Data[78:80], Bin)})
	data.List = append(data.List, Item{Key: "充电模式", Value: getParseData(xl.Data[80:82], Bin)})
	data.List = append(data.List, Item{Key: "最小功率", Value: getParseData(xl.Data[82:86], Bin)})
	data.List = append(data.List, Item{Key: "最大功率", Value: getParseData(xl.Data[86:90], Bin)})
	data.List = append(data.List, Item{Key: "最大允许电流", Value: getParseData(xl.Data[90:94], Bin)})
	data.List = append(data.List, Item{Key: "空载阈值电流", Value: getParseData(xl.Data[94:98], Bin)})
	data.List = append(data.List, Item{Key: "充满阈值电流", Value: getParseData(xl.Data[98:102], Bin)})
	data.List = append(data.List, Item{Key: "空载阈值功率", Value: getParseData(xl.Data[102:106], Bin)})
	data.List = append(data.List, Item{Key: "充满阈值功率", Value: getParseData(xl.Data[106:110], Bin)})
	data.List = append(data.List, Item{Key: "空载等待时间", Value: getParseData(xl.Data[110:114], Bin)})
	data.List = append(data.List, Item{Key: "充满等待时间", Value: getParseData(xl.Data[114:118], Bin)})
	data.List = append(data.List, Item{Key: "免费充电时间", Value: getParseData(xl.Data[118:122], Bin)})
	data.List = append(data.List, Item{Key: "最大允许充电时间", Value: getParseData(xl.Data[122:126], Bin)})
	data.List = append(data.List, Item{Key: "是否充满断电", Value: getParseData(xl.Data[126:128], Bin)})
	data.List = append(data.List, Item{Key: "功率分段1(单 价)", Value: getParseData(xl.Data[128:132], Bin)})
	data.List = append(data.List, Item{Key: "功率分段1（功 率）", Value: getParseData(xl.Data[132:136], Bin)})
	data.List = append(data.List, Item{Key: "功率分段2(单 价)", Value: getParseData(xl.Data[136:140], Bin)})
	data.List = append(data.List, Item{Key: "功率分段2（功 率）", Value: getParseData(xl.Data[140:144], Bin)})
	data.List = append(data.List, Item{Key: "功率分段3(单 价)", Value: getParseData(xl.Data[144:148], Bin)})
	data.List = append(data.List, Item{Key: "功率分段3（功 率）", Value: getParseData(xl.Data[148:152], Bin)})
	data.List = append(data.List, Item{Key: "功率分段4(单 价)", Value: getParseData(xl.Data[152:156], Bin)})
	data.List = append(data.List, Item{Key: "功率分段4（功 率）", Value: getParseData(xl.Data[156:160], Bin)})
	data.List = append(data.List, Item{Key: "功率分段5(单 价)", Value: getParseData(xl.Data[160:164], Bin)})
	data.List = append(data.List, Item{Key: "功率分段5（功 率）", Value: getParseData(xl.Data[164:168], Bin)})
	data.List = append(data.List, Item{Key: "功率分段6(单 价)", Value: getParseData(xl.Data[168:172], Bin)})
	data.List = append(data.List, Item{Key: "功率分段6（功 率）", Value: getParseData(xl.Data[172:176], Bin)})
	data.List = append(data.List, Item{Key: "功率分段7(单 价)", Value: getParseData(xl.Data[176:180], Bin)})
	data.List = append(data.List, Item{Key: "功率分段7（功 率）", Value: getParseData(xl.Data[180:184], Bin)})
	data.List = append(data.List, Item{Key: "功率分段8(单 价)", Value: getParseData(xl.Data[184:188], Bin)})
	data.List = append(data.List, Item{Key: "功率分段8（功 率）", Value: getParseData(xl.Data[188:192], Bin)})

	return
}

// XlBalanceModelSetting 交流桩计费模型设置
func (xl XlMessage) XlBalanceModelSetting() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "交流桩计费模型设置"
	// 解析数据
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[0:14], "0")})
	data.List = append(data.List, Item{Key: "计费模型编号", Value: getParseData(xl.Data[14:18], Bin)})
	data.List = append(data.List, Item{Key: "尖费电费费率", Value: getParseData(xl.Data[18:26], Bin)})
	data.List = append(data.List, Item{Key: "尖服务费费率", Value: getParseData(xl.Data[26:34], Bin)})
	data.List = append(data.List, Item{Key: "峰电费费率", Value: getParseData(xl.Data[34:42], Bin)})
	data.List = append(data.List, Item{Key: "峰服务费费率", Value: getParseData(xl.Data[42:50], Bin)})
	data.List = append(data.List, Item{Key: "平电费费率", Value: getParseData(xl.Data[50:58], Bin)})
	data.List = append(data.List, Item{Key: "平服务费费率", Value: getParseData(xl.Data[58:66], Bin)})
	data.List = append(data.List, Item{Key: "谷电费费率", Value: getParseData(xl.Data[66:74], Bin)})
	data.List = append(data.List, Item{Key: "谷服务费费率", Value: getParseData(xl.Data[74:82], Bin)})
	data.List = append(data.List, Item{Key: "计损比例", Value: getParseData(xl.Data[82:84], Bin)})
	data.List = append(data.List, Item{Key: "0：00～0：30时段费率号", Value: getParseData(xl.Data[84:86], Bin)})
	data.List = append(data.List, Item{Key: "0：30～1：00时段费率号", Value: getParseData(xl.Data[100:102], Bin)})
	data.List = append(data.List, Item{Key: "......", Value: getParseData(xl.Data[102:104], Bin)})
	data.List = append(data.List, Item{Key: "23：00～23：30时段费率号", Value: getParseData(xl.Data[104:106], Bin)})
	data.List = append(data.List, Item{Key: "23：30～0：00时段费率号", Value: getParseData(xl.Data[106:108], Bin)})

	return
}

// XlBusRealTimeData 汽车桩实时数据
func (xl XlMessage) XlBusRealTimeData() (data Data) {
	data.Cmd = xl.Command
	data.Remark = "汽车桩实时数据"
	// 解析数据
	data.List = append(data.List, Item{Key: "交易流水号", Value: strings.TrimLeft(xl.Data[0:32], "0")})
	data.List = append(data.List, Item{Key: "桩编号", Value: strings.TrimLeft(xl.Data[32:46], "0")})
	data.List = append(data.List, Item{Key: "枪号", Value: strings.TrimLeft(xl.Data[46:48], "0")})
	data.List = append(data.List, Item{Key: "状态", Value: getParseData(xl.Data[48:50], Bin)})
	data.List = append(data.List, Item{Key: "枪是否归位", Value: getParseData(xl.Data[50:52], Bin)})
	data.List = append(data.List, Item{Key: "是否插枪", Value: getParseData(xl.Data[52:54], Bin)})
	data.List = append(data.List, Item{Key: "输出电压", Value: getParseData(xl.Data[54:58], Bin)})
	data.List = append(data.List, Item{Key: "输出电流", Value: getParseData(xl.Data[58:62], Bin)})
	data.List = append(data.List, Item{Key: "枪线温度", Value: getParseData(xl.Data[62:64], Bin)})
	data.List = append(data.List, Item{Key: "枪线编码", Value: getParseData(xl.Data[64:80], Bin)})
	data.List = append(data.List, Item{Key: "SOC", Value: getParseData(xl.Data[80:82], Bin)})
	data.List = append(data.List, Item{Key: "电池组最高温度", Value: getParseData(xl.Data[82:84], Bin)})
	data.List = append(data.List, Item{Key: "累计充电时间", Value: getParseData(xl.Data[84:88], Bin)})
	data.List = append(data.List, Item{Key: "剩余时间", Value: getParseData(xl.Data[88:92], Bin)})
	data.List = append(data.List, Item{Key: "充电度数", Value: getParseData(xl.Data[92:100], Bin)})
	data.List = append(data.List, Item{Key: "计损充电度数", Value: getParseData(xl.Data[100:108], Bin)})
	data.List = append(data.List, Item{Key: "已充金额", Value: getParseData(xl.Data[108:116], Bin)})
	data.List = append(data.List, Item{Key: "系统故障", Value: getParseData(xl.Data[116:132], Bin)})
	data.List = append(data.List, Item{Key: "系统告警", Value: getParseData(xl.Data[132:148], Bin)})

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
	case Ascii:
		res = string(utils.StringToByte(str))
	}

	return
}
