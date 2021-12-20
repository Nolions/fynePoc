package view

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Nolions/fynePoc/pkg"
	"strconv"
)

func MainView(app *pkg.Application) {
	w := app.NewWindows("FynePoc")

	i := 0

	label := widget.NewLabel(strconv.Itoa(i))
	plusBtn := widget.NewButton("+", func() {
		i++
		label.SetText(strconv.Itoa(i))
	})

	reduceBtn := widget.NewButton("-", func() {
		if i > 0 {
			i--
			label.SetText(strconv.Itoa(i))
		}
	})
	newWinBtn := widget.NewButton("chose", func() {
		DialogDemo(app)
	})

	w.Win.SetContent(container.NewHBox(
		reduceBtn,
		label,
		plusBtn,
		newWinBtn,
		widget.NewButton("photo", func() {
			ChosePhoto(app)
		}),
	))

	w.Show()
}
