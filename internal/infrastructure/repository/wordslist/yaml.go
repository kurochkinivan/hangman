package wordslist

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type WordsList struct {
	Animals          CategoryData `yaml:"animals"`
	FruitsVegetables CategoryData `yaml:"fruits_vegetables"`
	Countries        CategoryData `yaml:"countries"`
}

type CategoryData struct {
	Easy   []Word `yaml:"easy"`
	Medium []Word `yaml:"medium"`
	Hard   []Word `yaml:"hard"`
}

type Word struct {
	Value string `yaml:"word"`
	Hint  string `yaml:"hint"`
}

func LoadWordsListFromYAML(pathToYAML string) (wl WordsList, err error) {
	if _, err := os.Stat(pathToYAML); err != nil {
		return WordsList{}, fmt.Errorf("failed to get file info: %w", err)
	}

	f, err := os.Open(pathToYAML)
	if err != nil {
		return WordsList{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	return loadWordsList(f)
}

// loadWordsList reads and decodes a YAML-formatted words list from the provided io.Reader.
// It returns the parsed WordsList and any error encountered during decoding.
//
// If the input is empty, an error indicating an empty file is returned.
// For other decoding errors, a descriptive error is returned.
func loadWordsList(r io.Reader) (WordsList, error) {
	var wl WordsList

	err := yaml.NewDecoder(r).Decode(&wl)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return WordsList{}, fmt.Errorf("file is empty: %w", err)
		}
		return WordsList{}, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return wl, nil
}
