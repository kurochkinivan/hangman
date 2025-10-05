package terminal

import (
	"bufio"
	"io"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameHandler struct {
	game     *entities.Game
	config   *entities.GameConfig
	showHint bool
	reader   *bufio.Reader
}

func NewGameHandler(wordsRepo entities.WordsRepository, r io.Reader) *GameHandler {
	return &GameHandler{
		game:     nil,
		config:   entities.NewGameConfig(wordsRepo),
		showHint: false,
		reader:   bufio.NewReader(r),
	}
}
