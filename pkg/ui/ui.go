package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	screen tcell.Screen
}

func NewScreen() *Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	return &Screen{s}
}

func (s *Screen) Close() {
	s.screen.Fini()
}

func (s *Screen) Clear() {
	s.screen.Clear()
}

func (s *Screen) Show() {
	s.screen.Show()
}

func (s *Screen) SetContent(x, y int, char rune, style tcell.Style) {
	s.screen.SetContent(x, y, char, nil, style) // Correct usage of SetContent
}

func (s *Screen) DrawSnippet(snippet []rune, input []rune, cursorPos int) {
	for i, char := range snippet {
		style := tcell.StyleDefault
		if i < len(input) {
			if input[i] == snippet[i] {
				style = style.Foreground(tcell.ColorGreen) // Correct character
			} else {
				style = style.Foreground(tcell.ColorRed) // Incorrect character
			}
		}
		s.SetContent(i, 0, char, style)
	}
	if cursorPos < len(snippet) {
		s.SetContent(cursorPos, 0, '|', tcell.StyleDefault) // Draw cursor
	}
}

func (s *Screen) DrawGameOver(snippet []rune) {
	gameOverText := "Game Over! Press ESC to exit."
	x := (len(snippet) - len(gameOverText)) / 2
	for i, char := range gameOverText {
		s.SetContent(x+i, 0, char, tcell.StyleDefault.Foreground(tcell.ColorYellow))
	}
}

func (s *Screen) InitialScreen() {
	s.Clear()

	title := "Welcome to 'howfast' a typing test cli game"
	instructions := "Press Enter to start | Press Esc to exit"
	subInstructions := "Type the snippet as fast and accurately you can"

	titleX := (s.width() - len(title)) / 2
	instructionsX := (s.width() - len(instructions)) / 2
	subInstructionsX := (s.width() - len(subInstructions)) / 2

	for i, char := range title {
		s.SetContent(titleX+i, 5, char, tcell.StyleDefault.Foreground(tcell.ColorYellow))
	}

	for i, char := range instructions {
		s.SetContent(instructionsX+i, 7, char, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	for i, char := range subInstructions {
		s.SetContent(subInstructionsX+i, 9, char, tcell.StyleDefault.Foreground(tcell.ColorGray))
	}

	s.Show()
}

func (s *Screen) width() int {
	w, _ := s.screen.Size()
	return w
}

func (s *Screen)WaitForStartOrExit()bool {
	for {
		ev := s.screen.PollEvent()
		switch ev := ev.(type){
			case *tcell.EventKey: 
			if ev.Key() == tcell.KeyEnter{
				return true
			}else if ev.Key() == tcell.KeyEsc {
				return false
			}
		}
	}
}