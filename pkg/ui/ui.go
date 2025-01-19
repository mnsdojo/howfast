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
		log.Fatalf("Error creating screen :%v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("Error initializing screen : %v", err)
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
	s.screen.SetContent(x, y, char, nil, style)
}

func (s *Screen) DrawSnippet(snippet []rune, input []rune, cursorPos int) {
	for i, char := range snippet {
		style := tcell.StyleDefault
		if i < len(input) {
			if input[i] == snippet[i] {
				style = style.Foreground(tcell.ColorGreen)
			} else {
				style = style.Foreground(tcell.ColorRed)
			}
		}
		s.SetContent(i, 0, char, style)

	}
	if cursorPos < len(snippet) {
		s.SetContent(cursorPos, 0, '|', tcell.StyleDefault)
	}
}

func (s *Screen) DrawGameOver(snippet []rune) {
	gameOverText := "Game Over! Press ESC to exit."
	x := (len(snippet) - len(gameOverText)) / 2
	for i, char := range gameOverText {
		s.SetContent(x+i, 0, char, tcell.StyleDefault.Foreground(tcell.ColorYellow))
	}
}
