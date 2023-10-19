package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/gaomengnan/fyne-demo/data"
)

func main() {
	confs := data.GetConfigs()

	app := app.New()

	// setting theme
	app.Settings().SetTheme(
		NewMyTheme(),
	)

	// set title
	window := app.NewWindow("Database Tool")
	window.Resize(fyne.NewSize(768, 480))
	window.CenterOnScreen()

	ui := newUI(window)
	// set content
	window.SetContent(
		ui.makeUI(),
	)

	// set main menu
	window.SetMainMenu(ui.makeMenu())

	// show create button
	if len(confs.Servers) == 0 {
		ui.showCreate(window)
	}

	// run
	window.ShowAndRun()
}
