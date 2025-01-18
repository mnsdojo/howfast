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
