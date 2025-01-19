package ui

import (
	"fmt"
	"log"
	"time"

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

func (s *Screen) DrawStats(timeTaken time.Duration, errors int, accuracy float64, wpm float64) {
	s.Clear()
	stats := fmt.Sprintf("Time: %v | Errors: %d | Accuracy: %.2f%% | WPM: %.2f", timeTaken, errors, accuracy, wpm)
	x := (s.width() - len(stats)) / 2
	for i, char := range stats {
		s.SetContent(x+i, 5, char, tcell.StyleDefault.Foreground(tcell.ColorYellow))
	}
	s.Show()
}
func (s *Screen) HandleTypingInput(snippetRunes []rune, input *[]rune, cursorPos *int) bool {
	ev := s.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			return true // Exit the game
		case tcell.KeyLeft:
			if *cursorPos > 0 {
				*cursorPos-- // Move cursor left
			}
		case tcell.KeyRight:
			if *cursorPos < len(*input) {
				*cursorPos++ // Move cursor right
			}
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			if *cursorPos > 0 {
				// Remove the character before the cursor
				*input = append((*input)[:*cursorPos-1], (*input)[*cursorPos:]...)
				*cursorPos-- // Move cursor left
			}
		case tcell.KeyRune:
			if *cursorPos < len(snippetRunes) {
				// Insert the typed character at the cursor position
				*input = append((*input)[:*cursorPos], append([]rune{ev.Rune()}, (*input)[*cursorPos:]...)...)
				*cursorPos++ // Move cursor right
			}
		}
	}

	// Check if the player has finished typing
	return len(*input) == len(snippetRunes)
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

func (s *Screen) WaitForStartOrExit() bool {
	for {
		ev := s.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				return true
			} else if ev.Key() == tcell.KeyEsc {
				return false
			}
		}
	}
}


func (s *Screen) WaitForRetryOrExit() bool {
	retryText := "Press R to retry | Press ESC to exit"
	x := (s.width() - len(retryText)) / 2

	for i, char := range retryText {
		s.SetContent(x+i, 7, char, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}
	s.Show()

	// Wait for the user to press R or ESC
	for {
		ev := s.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune && (ev.Rune() == 'r' || ev.Rune() == 'R') {
				return true // Retry the game
			} else if ev.Key() == tcell.KeyEscape {
				return false // Exit the program
			}
		}
	}
}
