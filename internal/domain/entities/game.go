package entities

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

type GameStatus int

const (
	GameStatusInProgress GameStatus = iota + 1
	GameStatusWon
	GameStatusLost
)

type Game struct {
	word           Word
	guessedLetters map[rune]bool
	maxAttempts    int
	wrongAttempts  int
}

func NewGame(word Word, maxAttempts int) *Game {
	return &Game{
		word:           word,
		guessedLetters: make(map[rune]bool),
		maxAttempts:    maxAttempts,
		wrongAttempts:  0,
	}
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
	remaining := g.maxAttempts - g.wrongAttempts
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (g *Game) Status() GameStatus {
	if g.wrongAttempts >= g.maxAttempts {
		return GameStatusLost
	}

	for _, char := range g.word.Value() {
		if char != ' ' && !g.IsLetterGuessed(char) {
			return GameStatusInProgress
		}
	}

	return GameStatusWon
}

func (g *Game) IsWon() bool {
	return g.Status() == GameStatusWon
}

func (g *Game) InProgress() bool {
	return g.Status() == GameStatusInProgress
}

func (g *Game) Word() Word {
	return g.word
}
