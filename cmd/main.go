package main

import (
	"github.com/Nolions/fynePoc/pkg"
	"github.com/Nolions/fynePoc/view"
)

func main() {
	a := pkg.NewApp()
	view.MainView(a)

	a.App.Run()
}
