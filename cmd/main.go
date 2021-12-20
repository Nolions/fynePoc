package main

import (
	"fynePoc/pkg"
	"fynePoc/view"
)

func main() {
	a := pkg.NewApp()
	view.MainView(a)

	a.App.Run()
}
