package pkg

import (
	"fyne.io/fyne/v2"
)

type Window struct {
	Win fyne.Window
}

func (w *Window) Show() {
	w.Win.Show()
}
