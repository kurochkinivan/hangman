package entities

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	apperr "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/lib/appErr"
)

type GameStatus int

const (
	GameStatusInProgress GameStatus = iota + 1
	GameStatusWon
	GameStatusLost
)

type Game struct {
	word           *Word
	config         *GameConfig
	guessedLetters map[rune]bool
	wrongAttempts  int
}

func NewGame(word *Word, config *GameConfig) (*Game, error) {
	if word == nil {
		return nil, apperr.NewAppErr("NewGame", "word cannot be nil")
	}

	if config == nil {
		return nil, apperr.NewAppErr("NewGame", "config cannot be nil")
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

func (g *Game) GuessedLetters() []rune {
	letters := make([]rune, 0, len(g.guessedLetters))

	for letter := range g.guessedLetters {
		letters = append(letters, letter)
	}

	slices.Sort(letters)

	return letters
}

func (g *Game) RemainingAttempts() int {
	remaining := g.config.MaxAttempts() - g.wrongAttempts
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (g *Game) Status() GameStatus {
	if g.wrongAttempts >= g.config.MaxAttempts() {
		return GameStatusLost
	}

	for _, char := range g.word.Value() {
		if char != ' ' && !g.IsLetterGuessed(char) {
			return GameStatusInProgress
		}
	}

	return GameStatusWon
}

func (g *Game) Word() *Word {
	return g.word
}

func (g *Game) Config() *GameConfig {
	return g.config
}

/*

Очевидно, что лучше вынести логику из infrastracture/terminal

Как это сделать?

Где-то должен создаться и ХРАНИТЬСЯ инстанс игры.

*/
