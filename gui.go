package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gaomengnan/fyne-demo/data"
	"github.com/gaomengnan/fyne-demo/database"
	"github.com/gaomengnan/fyne-demo/internal/dialogs"
)

type gui struct {
	w    fyne.Window
	left *widget.Tree

	// treedata
	treeData binding.StringTree
	// content                 *container.DocTabs
	currentActiveConnection *string

	// tabs
	tabs map[string]*container.DocTabs
}

func newUI(win fyne.Window) *gui {
	return &gui{
		w: win,
		tabs: map[string]*container.DocTabs{
			"default": container.NewDocTabs(container.NewTabItem("Sql Query", widget.NewLabel("Empty"))),
		},
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
	left := g.makeLeftContent()
	right := widget.NewLabel("Right")

	// content
	// content := widget.NewLabel("Content")
	// content.Alignment = fyne.TextAlignCenter
	// content := canvas.NewRectangle(
	// 	color.Gray{
	// 		Y: 0xee,
	// 	},
	// )

	content := g.tabs["default"]

	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}

	objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}

	g.left = left

	return container.New(newLayout(top, left, right, content, dividers), objs...)
}

func (g *gui) makeMenu() *fyne.MainMenu {
	newConnectionMenu := fyne.NewMenu(
		"Edit",
		fyne.NewMenuItem(
			"New",
			func() {},
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
		seriallize := entry.Get()
		err := database.Connect(seriallize)
		if err != nil {
			dialog.ShowError(err, g.w)
			return
		}
		err = entry.Save()
		if err != nil {
			dialog.ShowError(err, g.w)
			return
		}
		err = g.treeData.Append(binding.DataTreeRootID, seriallize.DSN(), seriallize.Name)
		if err != nil {
			dialog.ShowError(err, g.w)
			return
		}
		dialog.ShowInformation(seriallize.DSN(), "Connect Successfully", g.w)
		wizard.Hide()
	}

	// 创建额外的按钮
	extraButton := widget.NewButton("Test Connection", func() {
		// 在这里添加按钮点击事件的处理逻辑
		if err := form.Validate(); err != nil {
			return
		}
		err := database.TestConnect(entry.Get())
		if err != nil {
			dialog.ShowError(err, g.w)
			return
		}
		dialog.ShowInformation(entry.Get().DSN(), "Connect Successfully", g.w)
	})

	// 创建一个容器，包括表单和额外按钮
	content := container.NewVBox(
		extraButton,
		form,
	)
	return content
}

func (g *gui) makeLeftContent() *widget.Tree {
	// 创建主级折叠面板
	treeData := binding.NewStringTree()
	for _, v := range data.GlobalConfigure.Servers {
		err := treeData.Append(binding.DataTreeRootID, v.DSN(), v.Name)
		if err != nil {
			dialog.ShowError(err, g.w)
			return nil
		}
	}
	mainAccordion := widget.NewTreeWithData(treeData, func(branch bool) fyne.CanvasObject {
		label := widget.NewLabel("text")
		container := container.NewGridWithColumns(1, label)
		return container
	},
		func(di binding.DataItem, b bool, co fyne.CanvasObject) {
			l := co.(*fyne.Container)
			label := l.Objects[0].(*widget.Label)
			u, _ := di.(binding.String).Get()
			label.SetText(u)
		})

	mainAccordion.OnSelected = func(uid widget.TreeNodeID) {
		val, _ := treeData.GetValue(uid)
		var connectData *data.SerializationConnectionData
		for _, v := range data.GlobalConfigure.Servers {
			if v.Name == val {
				connectData = v
				break
			}

		}
		if connectData.Connection == nil {
			err := database.Connect(connectData)
			if err != nil {
				dialog.ShowError(err, g.w)
				return
			}
		}
		//
		fmt.Println(connectData)
		rows, err := connectData.Connection.Query("SHOW DATABASES")
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var dbName string
			err := rows.Scan(&dbName)
			if err != nil {
				panic(err)
			}

			err = treeData.Append(uid, fmt.Sprintf("%s-%s", uid, dbName), dbName)
			if err != nil {
				return
			}
		}
	}

	g.treeData = treeData
	return mainAccordion
}
