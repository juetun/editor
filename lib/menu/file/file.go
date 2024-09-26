package file

import "fyne.io/fyne/v2"

func GetFile(app fyne.App, topWindow fyne.Window) (mainMenu *fyne.Menu) {
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	file.Items = append(file.Items, aboutItem)
	return
}
