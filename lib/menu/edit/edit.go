package edit

import "fyne.io/fyne/v2"

func GetEdit(app fyne.App, topWindow fyne.Window) (mainMenu *fyne.Menu) {
	mainMenu = fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem)
	return
}
