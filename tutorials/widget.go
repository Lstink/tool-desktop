package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"lstink.github.com/lstink/tool-desktop/internal/device"
	"lstink.github.com/lstink/tool-desktop/internal/message"
)

func makeFormTabXl(w fyne.Window) fyne.CanvasObject {
	var msg = binding.NewString()
	// 下面是table数据
	input := widget.NewEntryWithData(msg)
	input.PlaceHolder = "请输入报文内容"
	input.Validator = validation.NewRegexp(`68[\d|\w]{5,}`, "无效的报文")

	form := &widget.Form{
		OnCancel: func() {
			msg.Set("")
		},
		OnSubmit: func() {
			defer func() {
				if err := recover(); err != nil { //产生了panic异常
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Content: "数据格式不正确",
					})
				}
			}()

			w2 := fyne.CurrentApp().NewWindow("解析结果")
			table := handleTableData(w, input.Text, device.XLDevice)
			w2.SetContent(table)
			w2.Show()
			// 解析数据
		},
	}

	form.Append("Message", input)

	return container.NewVBox(form)
}

func makeFormTabCx(w fyne.Window) fyne.CanvasObject {
	var msg = binding.NewString()
	// 下面是table数据
	input := widget.NewEntryWithData(msg)
	input.PlaceHolder = "请输入报文内容"
	input.Validator = validation.NewRegexp(`68[\d|\w]{5,}`, "无效的报文")

	form := &widget.Form{
		OnCancel: func() {
			msg.Set("")
		},
		OnSubmit: func() {
			defer func() {
				if err := recover(); err != nil { //产生了panic异常
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Content: "数据格式不正确",
					})
				}
			}()

			w2 := fyne.CurrentApp().NewWindow("解析结果")
			table := handleTableData(w, input.Text, device.XLDevice)
			w2.SetContent(table)
			w2.Show()
			// 解析数据
		},
	}

	form.Append("Message", input)

	return container.NewVBox(form)
}

// 解析数据
func handleTableData(_ fyne.Window, msg string, flag device.DeviceType) fyne.CanvasObject {

	// 解析获取到的数据
	var res = binding.NewString()
	var data []message.Data
	var errMsg string

	if msg != "" {
		// 解析数据
		data, errMsg = device.ParseDataForMessage(msg, flag)
		if errMsg != "" {
			res.Set(errMsg)
		}
	}

	t := getTable(data)

	res.Set(fmt.Sprintf("数据解析成功！"))

	c := container.NewVBox(widget.NewLabelWithData(res), container.New(layout.NewGridWrapLayout(fyne.NewSize(640, 460)), t))
	return c

}

func getTable(data []message.Data) fyne.CanvasObject {
	t := widget.NewTable(
		// 数据列数量， 数据行数量
		func() (int, int) {
			return len(data) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				switch id.Col {
				case 0:
					label.SetText("序号")
				case 1:
					label.SetText("字段")
				default:
					label.SetText("值")
				}
			} else {
				switch id.Col {
				case 0:
					label.SetText(fmt.Sprintf("%d", id.Row))
				case 1:
					label.SetText(data[id.Row-1].Key)
				default:
					label.SetText(data[id.Row-1].Value)
				}
			}

		})
	t.SetColumnWidth(0, 120)
	t.SetColumnWidth(1, 200)

	return t
}
