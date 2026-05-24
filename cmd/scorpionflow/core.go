package main

import (
	_ "embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/icon.png
var icon []byte

func main() {
	a := app.New()
	w := a.NewWindow("ScorpionFlow")

	// Label, das später geändert werden kann
	output := widget.NewLabel("Noch keine Nachricht")

	// Bild
	img := canvas.NewImageFromResource(fyne.NewStaticResource("icon.png", icon))
	img.FillMode = canvas.ImageFillContain

	// Button
	btn := widget.NewButton("Drück mich", func() {
		fmt.Println("Es wurde ein Knopf gedrückt!")
	})

	// Layout
	w.SetContent(container.NewVBox(
		widget.NewLabel("Hallo Welt"),
		img,
		btn,
		output,
	))

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
