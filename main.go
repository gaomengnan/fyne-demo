package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	// setting theme
	app.Settings().SetTheme(
		NewMyTheme(),
	)

	// set title
	window := app.NewWindow("Database Tool")

	// set content
	window.SetContent(
		MakeUI(),
	)

	// run
	window.ShowAndRun()
}
