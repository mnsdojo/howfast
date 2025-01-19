package main

import (
	"image/draw"
	"time"

	"github.com/mnsdojo/howfast/pkg/ui"
	"github.com/mnsdojo/howfast/pkg/utils"
)

func main() {
	screen := ui.NewScreen()
	defer screen.Close()
	screen.InitialScreen()
	if !screen.WaitForStartOrExit(){
		return
	}
	snippet,err := utils.GetRandomSnippet()
	if err!=nil {
		return
	}
	snippetRunes := []rune(snippet)
	var input[]rune 
	cursorPos := 0 
	startTime := time.Now()
	for {
		screen.Clear()
		screen.DrawSnippet(snippetRunes, input, cursorPos)
		screen.Show()
		
		
	}
}
