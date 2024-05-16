package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyLightTheme struct{}

var _ fyne.Theme = (*MyLightTheme)(nil)

// return bundled font resource
func (*MyLightTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.LightTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.LightTheme().Font(s)
		}
		return resource1Ttf
	}
	if s.Italic {
		return theme.LightTheme().Font(s)
	}
	return resource1Ttf
}

func (*MyLightTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.LightTheme().Color(n, v)
}

func (*MyLightTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.LightTheme().Icon(n)
}

func (*MyLightTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.LightTheme().Size(n)
}
