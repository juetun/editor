package help

import "fyne.io/fyne/v2"

func GetHelp(app fyne.App, topWindow fyne.Window) (mainMenu *fyne.Menu) {
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, topWindow)
	})
	mainMenu = fyne.NewMenu("Help", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem)
	return
}
