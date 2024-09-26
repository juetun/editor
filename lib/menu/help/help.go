package help

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/editor/lib/common"
)

func GetHelp(app fyne.App, topWindow fyne.Window) (mainMenu *fyne.Menu) {
	//helpMenu := fyne.NewMenu("Help15",
	//	fyne.NewMenuItem("Documentation", func() {
	//		u, _ := url.Parse("https://developer.fyne.io")
	//		_ = app.OpenURL(u)
	//	}),
	//	fyne.NewMenuItem("Support16", func() {
	//		u, _ := url.Parse("https://fyne.io/support/")
	//		_ = app.OpenURL(u)
	//	}),
	//	fyne.NewMenuItemSeparator(),
	//	fyne.NewMenuItem("Sponsor17", func() {
	//		u, _ := url.Parse("https://fyne.io/sponsor/")
	//		_ = app.OpenURL(u)
	//	}))

	performFind := func() { fmt.Println("Menu Find13") }
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: topWindow.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		cutShortcut := &fyne.ShortcutCut{Clipboard: topWindow.Clipboard()}
		common.ShortcutFocused(cutShortcut, topWindow)
	})
	copyItem := fyne.NewMenuItem("Copy", func() {
		copyShortcut := &fyne.ShortcutCopy{Clipboard: topWindow.Clipboard()}
		common.ShortcutFocused(copyShortcut, topWindow)
	})
	pasteItem := fyne.NewMenuItem("Paste", func() {
		common.ShortcutFocused(pasteShortcut, topWindow)
	})
	findItem := fyne.NewMenuItem("Find14", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	topWindow.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	pasteItem.Shortcut = pasteShortcut
	mainMenu = fyne.NewMenu("Help", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem)
	return
}
