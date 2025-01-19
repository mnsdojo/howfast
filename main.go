package main

import "github.com/mnsdojo/howfast/pkg/ui"

func main() {
	screen := ui.NewScreen()
	defer screen.Close()
	screen.InitialScreen()
}
