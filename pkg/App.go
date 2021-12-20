package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"sync"
)

type Application struct {
	App fyne.App
}

var once sync.Once
var instance *Application

func NewApp() *Application {
	once.Do(func() {
		instance = &Application{
			App: app.New(),
		}
	})

	return instance
}

func (app *Application) NewWindows(title string) *Window {
	return &Window{
		Win: app.App.NewWindow(title),
	}
}
