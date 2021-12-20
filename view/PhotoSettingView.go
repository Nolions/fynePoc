package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Nolions/fynePoc/pkg"
)

//var im *canvas.Image

func ChosePhoto(app *pkg.Application) {
	w := app.NewWindows("chose photo")

	im := pkg.NewImageView()

	d := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
		if closer != nil {
			//im.View.File = closer.URI().Path()
			im.Refresh(closer.URI().Path())
		}
	}, w.Win)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".png", "jpg"}))

	b := widget.NewButton("Chose File", func() {
		d.Show()
	})

	w.Win.SetContent(container.NewVBox(
		im.View,
		b,
	))
	w.Show()
}
