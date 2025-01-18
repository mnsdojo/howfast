package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

const (
	path = "assets/code.json"
)

func calculateErrors(snippetChars, typedChars []rune, cursorPos int) int {

	errors := 0
	for i := 0; i < cursorPos; i++ {
		if typedChars[i] != snippetChars[i] {
			errors++
		}

	}
	return errors
}

func calculateAccuracy(snippetChars, typedChars []rune, cursorPos int) float64 {
	correctChars := 0
	for i := 0; i < cursorPos; i++ {
		if typedChars[i] == snippetChars[i] {
			correctChars++
		}

	}
	return float64(correctChars) / float64(len(snippetChars)) * 100

}
func calculateWPM(snippetChars []rune, timeTaken time.Duration)float64 {
	words := float64(len(snippetChars))/5.0
	minutes := timeTaken.Minutes()
	return words/minutes
}
func getRandomSnippet() (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()
	var data struct {
		Paragraphs []string `json:"paragraphs"`
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return "", err
	}
	if len(data.Paragraphs) == 0 {
		return "", errors.New("no paragraphs found in the file ")
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(data.Paragraphs))
	randomParagraph := data.Paragraphs[randomIndex]
	return randomParagraph, nil
}
