package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getIconPath() string {
	// Versuche verschiedene mögliche Pfade
	paths := []string{
		"cmd/scorpionflow/assets/icon.png",
		"assets/icon.png",
		"./assets/icon.png",
		"/usr/share/icons/hicolor/256x256/apps/scorpionflow.png",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

func main() {
	a := app.New()
	a.SetIcon(fyne.NewStaticResource("icon.png", []byte{}))

	// Icon Pfad ermitteln
	iconPath := getIconPath()
	fmt.Println("Icon Pfad:", iconPath)

	var iconData []byte
	if iconPath != "" {
		var err error
		iconData, err = os.ReadFile(iconPath)
		if err == nil && len(iconData) > 0 {
			fmt.Printf("Icon geladen: %d bytes\n", len(iconData))
			// Icon mit korrektem Namen registrieren
			res := fyne.NewStaticResource("icon.png", iconData)
			a.SetIcon(res)
			fmt.Println("Icon auf App gesetzt")
		} else if err != nil {
			fmt.Printf("Fehler beim Laden des Icons: %v\n", err)
		}
	}

	w := a.NewWindow("ScorpionFlow")

	// Auch auf dem Fenster setzen
	if len(iconData) > 0 {
		res := fyne.NewStaticResource("icon.png", iconData)
		w.SetIcon(res)
		fmt.Println("Icon auf Fenster gesetzt")
	}

	// Label, das später geändert werden kann
	output := widget.NewLabel("Noch keine Nachricht")

	// Bild anzeigen
	var img *canvas.Image
	if len(iconData) > 0 {
		res := fyne.NewStaticResource("icon.png", iconData)
		img = canvas.NewImageFromResource(res)
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(200, 200))
	}

	// Button
	btn := widget.NewButton("Drück mich", func() {
		fmt.Println("Es wurde ein Knopf gedrückt!")
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Hallo Welt"),
	)
	if img != nil {
		content.Add(img)
	}
	content.Add(btn)
	content.Add(output)

	w.SetContent(content)

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Fehler beim Laden von %s: %v\n", path, err)
		return []byte{}
	}
	fmt.Printf("Datei %s geladen: %d bytes\n", path, len(data))
	return data
}
