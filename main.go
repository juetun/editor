// Package main provides various examples of Fyne API capabilities.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/cmd/fyne_demo/tutorials"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/editor/lib/common"
	"github.com/editor/lib/menu"
	"github.com/editor/lib/params"
)

var topWindow fyne.Window

func main() {
	params.LoadConfig("app.yaml") //加载配置文件
	app := app.NewWithID(params.AppInfo.AppId)
	app.SetIcon(data.FyneLogo)
	common.MakeTray(app, params.Io)
	common.LogLifecycle(app, params.Io)
	window := app.NewWindow("")
	menu := menu.GetMenu(app, window)
	window.SetMainMenu(menu)
	window.SetMaster()
	content := container.NewStack()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	topWindow = window
	setTutorial := func(t tutorials.Tutorial) {
		if fyne.CurrentDevice().IsMobile() {
			child := app.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = window
			})
			return
		}
		title.SetText(t.Title)
		intro.SetText(t.Intro)

		if t.Title == "Welcome" {
			title.Hide()
			intro.Hide()
		} else {
			title.Show()
			intro.Show()
		}
		content.Objects = []fyne.CanvasObject{t.View(window)}
		content.Refresh()
	}
	if fyne.CurrentDevice().IsMobile() {
		window.SetContent(common.MakeNav(setTutorial, false))
	} else {
		split := container.NewHSplit(common.MakeNav(setTutorial, true), container.NewBorder(
			container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content))
		split.Offset = 0
		window.SetContent(split)
	}
	window.Resize(fyne.NewSize(640, 460))
	window.ShowAndRun()
}
