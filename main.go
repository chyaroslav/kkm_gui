package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type kkm struct {
	App fyne.App
	Mw  fyne.Window
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Кассовое место")
	k := kkm{App: myApp, Mw: myWindow}

	k.createLayout()
	k.StartApp()

	//myWindow.SetContent(content)

}
