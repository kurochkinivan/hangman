package simulate

import (
	"fmt"
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
	if utf8.RuneCountInString(gh.game.Word().Value()) != utf8.RuneCountInString(gh.guessed) {
		fmt.Println("Lengths of given word and guessed word do not match")
		return
	}

	for _, r := range gh.guessed {
		gh.game.GuessLetter(r)
	}

	if gh.game.IsWon() {
		fmt.Printf("%s;POS\n", gh.game.WordMask())
	} else {
		fmt.Printf("%s;NEG\n", gh.game.WordMask())
	}
}
