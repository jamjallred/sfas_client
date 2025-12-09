package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	base fyne.Theme
}

func newMyTheme() fyne.Theme {
	return &myTheme{base: theme.DefaultTheme()}
}

func (t *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	lightVariant := theme.VariantDark
	switch name {
	case theme.ColorNameForeground:
		return color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	}

	return t.base.Color(name, lightVariant)
}

func (t *myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.base.Font(style)
}

func (t *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.base.Icon(name)
}

func (t *myTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return theme.DefaultTheme().Size(name) * 1.2
	}
	return t.base.Size(name)
}
