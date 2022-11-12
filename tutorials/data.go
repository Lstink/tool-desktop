package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"xl": {"小蓝协议",
			"输入报文内容，点击解析",
			makeFormTabXl,
			true,
		},
		"cx": {"持续协议",
			"输入报文内容，点击解析",
			makeFormTabCx,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"": {"xl", "cx"},
	}
)
