package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/theme"
  "image/color"
)

type dark struct {}

var _ fyne.Theme = (*dark)(nil)

func (m dark) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
  switch name {
  case theme.ColorNameBackground:
    return color.NRGBA{R: 11, G:11, B:9, A:1}
  case theme.ColorNameForeground:
    return color.White
  case theme.ColorNamePrimary:
    return color.NRGBA{R:197, G:235, B:195, A:1}
  default:
  //
}
  
  return theme.DefaultTheme().Color(name, variant)
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
