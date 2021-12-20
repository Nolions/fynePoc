package view

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"fynePoc/internal"
	"fynePoc/pkg"
)

func DialogDemo(app *pkg.Application) {
	w := app.NewWindows("Dialog")

	w.Win.Resize(fyne.NewSize(200, 200))

	d1 := dialog.NewConfirm("title", "message", func(b bool) {
		fmt.Printf("b:%+v\n", b)
	}, w.Win)

	d2 := dialog.NewError(internal.DemoError{
		Title:   "error title",
		Message: "error message",
	}, w.Win)

	d3 := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
		if closer == nil {
			println("cancel")
		}
	}, w.Win)
	d3.SetFilter(storage.NewExtensionFileFilter([]string{".png", "jpg"}))

	d4 := dialog.NewInformation("title", "message", w.Win)

	w.Win.SetContent(container.NewHBox(
		widget.NewButton("close", func() {
			w.Win.Close()
		}),
		widget.NewButton("ConfirmDialog", func() {
			d1.Show()
		}),
		widget.NewButton("ErrorDialog", func() {
			d2.Show()
		}),
		widget.NewButton("Chose File", func() {
			d3.Show()
		}),

		widget.NewButton("Info", func() {
			d4.Show()
		}),
	))

	w.Show()
}
