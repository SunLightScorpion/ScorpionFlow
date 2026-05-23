package runtime

import "fmt"

type App struct {
}

func NewApp() *App {
	fmt.Println("New App")
	return &App{}
}

func (a *App) Test() {
	fmt.Println("Test")
}
