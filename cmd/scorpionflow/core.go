package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("ScorpionFlow")

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(widget.NewLabel("Hallo Welt"))
	w.ShowAndRun()
}
