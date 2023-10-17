package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NameBanner() fyne.CanvasObject {
	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.HomeIcon(),
			func() {},
		),
	)
	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain
	return container.NewStack(toolBar, container.NewPadded(logo))

}

func MakeUI() fyne.CanvasObject {
	top := NameBanner()
	left := widget.NewLabel("Left")
	right := widget.NewLabel("Right")

	// content
	// content := widget.NewLabel("Content")
	// content.Alignment = fyne.TextAlignCenter
	content := canvas.NewRectangle(
		color.Gray{
			Y: 0xee,
		},
	)

	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}

	objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}

	return container.New(newLayout(top, left, right, content, dividers), objs...)
}
