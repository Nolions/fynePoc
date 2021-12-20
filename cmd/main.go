package main

import (
	"github.com/Nolions/fynePoc/pkg"
	"github.com/Nolions/fynePoc/view"
	"os"
)

func main() {
	err := os.Setenv("FYNE_FONT", "./assets/wqy-microhei.ttc")
	if err != nil {
		return
	}
	a := pkg.NewApp()
	view.MainView(a)

	a.App.Run()

	err = os.Unsetenv("FYNE_FONT")
	if err != nil {
		return
	}
}
