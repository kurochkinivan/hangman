package terminal

import (
	"bufio"
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameHandler struct {
	game     Game
	config   Config
	showHint bool
	reader   *bufio.Reader
	out      io.Writer
}

func NewGameHandler(wordsRepo entities.WordsRepository, r io.Reader, out io.Writer) *GameHandler {
	return &GameHandler{
		game:     nil,
		config:   entities.NewGameConfig(wordsRepo),
		showHint: false,
		reader:   bufio.NewReader(r),
		out:      out,
	}
}

type Game interface {
	IsLetterGuessed(r rune) bool
	GuessLetter(r rune) bool
	InProgress() bool
	IsWon() bool
	Word() entities.Word

	GuessedLetters() []rune
	WordMask() string
	RemainingAttempts() int
}

type Config interface {
	GenerateWord() (entities.Word, error)
	SetLevel(level entities.Level)
	SetCategory(category entities.Category)

	Level() entities.Level
	Category() entities.Category
}
