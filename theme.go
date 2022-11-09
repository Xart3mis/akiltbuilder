package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dark struct{}

var _ fyne.Theme = (*dark)(nil)

func (m dark) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 4, G: 4, B: 4, A: 1}
	case theme.ColorNameForeground:
    return color.RGBA{R:249, G:244, B:245, A:1}
  // return color.White
	case theme.ColorNamePrimary:
		return color.RGBA{R: 113, G: 128, B: 172, A: 1}
  case theme.ColorNameFocus:
    return color.RGBA{R:28, G:33, B:29, A:1}

	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m dark) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m dark) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m dark) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
