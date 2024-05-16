package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyDarkTheme struct{}

var _ fyne.Theme = (*MyDarkTheme)(nil)

// return bundled font resource
func (*MyDarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DarkTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DarkTheme().Font(s)
		}
		return resource1Ttf
	}
	if s.Italic {
		return theme.DarkTheme().Font(s)
	}
	return resource1Ttf
}

func (*MyDarkTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DarkTheme().Color(n, v)
}

func (*MyDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DarkTheme().Icon(n)
}

func (*MyDarkTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DarkTheme().Size(n)
}
