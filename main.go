package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	file, err := os.ReadFile("./Icon.png")
	if err != nil {
		return
	}
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	myWindow.Resize(fyne.NewSize(500, 500))

	// btn container
	btn := widget.NewButtonWithIcon("点击我", fyne.NewStaticResource("logo", file), func() {
		fmt.Println("press")
	})
	btnContainer := container.NewCenter(btn)
	btnContainer.Resize(fyne.NewSize(50, 50))
	btn.Resize(fyne.NewSize(50, 50))

	myWindow.SetContent(btnContainer)

	myWindow.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
