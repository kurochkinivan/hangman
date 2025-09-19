package terminal

import (
	"bufio"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
)

type GameHandler struct {
	gameUseCase GameUseCase
	showHint    bool
	reader      *bufio.Reader
}

type GameUseCase interface {
	StartNewGame() error
	InProgress() bool
	GuessLetter(input string) (dto.UserMessage, error)
	GameInfo() dto.GameInfo
	GameResult() dto.GameResult
	Hint() string

	SetUpLevel(id int) error
	SetUpCategory(id int) error
	Categories() []dto.Category
	Levels() []dto.Level
}

func NewGameHandler(gameUseCase GameUseCase) *GameHandler {
	return &GameHandler{
		gameUseCase: gameUseCase,
		showHint:    false,
		reader:      bufio.NewReader(os.Stdin),
	}
}
