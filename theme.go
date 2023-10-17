//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	fyne.Theme
}

// func (m *myTheme) Font(_ fyne.TextStyle) fyne.Resource {
// 	panic("not implemented") // TODO: Implement
// }

// func (m *myTheme) Icon(_ fyne.ThemeIconName) fyne.Resource {
// 	panic("not implemented") // TODO: Implement
// }

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return m.Theme.Size(name)
}

func (t *myTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func NewMyTheme() fyne.Theme {
	return &myTheme{
		Theme: theme.DefaultTheme(),
	}
}
