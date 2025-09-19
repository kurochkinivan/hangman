package terminal

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
)

func (gh *GameHandler) displayGameState(info dto.GameInfo) {
	gh.clearTerminal()

	fmt.Println(hangStates[info.RemainingAttempts])
	fmt.Printf("Level: %s, Category: %s\n", info.Level, info.Category)
	fmt.Printf("Word: %s\n", info.WordMask)
	fmt.Printf("Remaining attempts: %d\n", info.RemainingAttempts)

	gh.printHint()

	fmt.Print("Guessed letters: ")
	for _, letter := range info.GuessedLetters {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}

func (gh *GameHandler) displayGameResult(grs dto.GameResult) {
	if grs.IsWon {
		fmt.Printf("🎉 Congratulations! You won! The word was: %s\n", grs.Word)
	} else {
		fmt.Printf("💀 Game over! The word was: %s\n", grs.Word)
	}
}

func (gh *GameHandler) printHint() {
	if gh.showHint {
		fmt.Println("Hint:", gh.gameUseCase.Hint())
	}
}

func (gh *GameHandler) printMenu() {
	fmt.Println("[1] Start Game")
	fmt.Println("[2] Settings")
	fmt.Println("[3] Exit")
}

func (gh *GameHandler) printError(err error, context string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to %s: %v\n", context, err)
	}
}
