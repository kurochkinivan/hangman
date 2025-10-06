package simulate

import (
	"fmt"
	"os"
	"unicode/utf8"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameHandler struct {
	game    *entities.Game
	guessed string
}

func NewGameHandler(word string, guessed string) *GameHandler {
	wrd := entities.NewWord(word, "")
	maxAttempts := len(guessed) + 1
	game := entities.NewGame(wrd, maxAttempts)

	return &GameHandler{
		game:    game,
		guessed: guessed,
	}
}

func (gh *GameHandler) Start() {
	result, err := gh.simulateGame()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to simulate game: %v", err)
	}

	if result.IsWon() {
		fmt.Printf("%s;POS\n", result.WordMask())
	} else {
		fmt.Printf("%s;NEG\n", result.WordMask())
	}
}

func (gh *GameHandler) simulateGame() (entities.GameResult, error) {
	if utf8.RuneCountInString(gh.game.Word().Value()) != utf8.RuneCountInString(gh.guessed) {
		return entities.GameResult{}, fmt.Errorf("lengths of given word and guessed word do not match")
	}

	for _, r := range gh.guessed {
		gh.game.GuessLetter(r)
	}

	mask := gh.game.WordMask()
	isWon := gh.game.IsWon()

	return entities.NewGameResult(mask, isWon), nil
}
