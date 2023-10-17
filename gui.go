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
	return container.NewStack(toolBar, logo)

}

func MakeUI() fyne.CanvasObject {
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

	return container.NewBorder(
		NameBanner(),
		nil,
		left,
		right,
		content,
	)
}
