package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makeTableTab(_ fyne.Window) fyne.CanvasObject {
	t := widget.NewTable(
		// 数据列数量， 数据行数量
		func() (int, int) { return 20, 3 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText("A longer cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
			}
		})
	t.SetColumnWidth(0, 34)
	t.SetColumnWidth(1, 102)
	return t
}
