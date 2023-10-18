package main

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gaomengnan/fyne-demo/data"
	"github.com/gaomengnan/fyne-demo/internal/dialogs"
)

type gui struct {
	w fyne.Window
}

func newUI(win fyne.Window) *gui {
	return &gui{
		w: win,
	}

}
func (g *gui) makeBanner() fyne.CanvasObject {
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

func (g *gui) makeUI() fyne.CanvasObject {
	top := g.makeBanner()
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

func (g *gui) makeMenu() *fyne.MainMenu {
	newConnectionMenu := fyne.NewMenu(
		"Edit",
		fyne.NewMenuItem(
			"New",
			g.openCreateConnection,
		),
	)
	return fyne.NewMainMenu(newConnectionMenu)
}

func (g *gui) showCreate(w fyne.Window) {
	var wizard *dialogs.Wizard
	open := widget.NewButton("New", func() {
		wizard.Push("Connection Details", g.makeCreate(wizard))
	})
	buttons := container.NewGridWithColumns(1, open)
	mainPage := container.NewVBox(buttons)
	wizard = dialogs.NewWizard("Create Connection", mainPage)
	wizard.Show(w)
	wizard.Resize(mainPage.MinSize().AddWidthHeight(300, 40))
}
func (g *gui) makeCreate(wizard *dialogs.Wizard) fyne.CanvasObject {
	entry := data.NewConnectionData()
	form := widget.NewForm(
		widget.NewFormItem("Name", entry.Name),
		widget.NewFormItem("Host", entry.Host),
		widget.NewFormItem("Port", entry.Port),
		widget.NewFormItem("UserName", entry.User),
		widget.NewFormItem("Password", entry.Password),
	)

	form.OnSubmit = func() {
		if entry.Name.Text == "" {
			dialog.ShowError(errors.New("Empty Name"), g.w)
			return
		}
	}
	return form
}

func (g *gui) openCreateConnection() {
	entry := data.NewConnectionData()
	testButton := widget.NewButton("Test", func() {
		// 处理提交操作
	})
	items := []*widget.FormItem{
		{
			Text:   "Name:",
			Widget: entry.Name,
		},
		{
			Text:   "Host:",
			Widget: entry.Host,
		},
		{
			Text:   "Port:",
			Widget: entry.Port,
		},
		{
			Text:   "User:",
			Widget: entry.User,
		},
		{
			Text:   "Password:",
			Widget: entry.Password,
		},
		widget.NewFormItem("Test Connection", testButton),
	}
	dialog.ShowForm("NewConnection", "Submit", "Cancle", items, func(b bool) {
		if b {
			// name := entry.Name.Text
			entry.Save()
		}
	}, g.w)
}
