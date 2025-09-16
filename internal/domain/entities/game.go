package entities

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Game struct {
	word           *Word
	config         *GameConfig
	guessedLetters map[rune]bool
	wrongAttempts  int
}

func NewGame(word *Word, config *GameConfig) (*Game, error) {
	if word == nil {
		return nil, errors.New("word cannot be nil")
	}

	if config == nil {
		return nil, errors.New("config cannot be nil")
	}

	return &Game{
		word:           word,
		config:         config,
		guessedLetters: make(map[rune]bool),
		wrongAttempts:  0,
	}, nil
}

func (g *Game) GuessLetter(r rune) bool {
	r = unicode.ToLower(r)

	if g.IsLetterGuessed(r) {
		return false
	}

	g.guessedLetters[r] = true

	if g.word.Contains(r) {
		return true
	}

	g.wrongAttempts++
	return false
}

func (g *Game) WordMask() string {
	var b strings.Builder
	b.Grow(utf8.RuneCountInString(g.word.Value()))

	for _, char := range g.word.Value() {
		if char == ' ' {
			b.WriteRune(' ')
		} else if g.IsLetterGuessed(char) {
			b.WriteRune(char)
		} else {
			b.WriteRune('*')
		}
	}

	return b.String()
}

func (g *Game) IsLetterGuessed(r rune) bool {
	_, used := g.guessedLetters[r]
	return used
}
