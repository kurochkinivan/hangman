package terminal

import (
	"fmt"
	"os"
)

func (gh *GameHandler) displayGameState() {
	gh.clearTerminal()
	remaining := gh.game.RemainingAttempts()

	fmt.Fprintln(gh.out, hangStates[len(hangStates)-remaining])
	fmt.Fprintf(gh.out, "Level: %s, Category: %s\n", gh.config.Level(), gh.config.Category())
	fmt.Fprintf(gh.out, "Word: %s\n", gh.game.WordMask())
	fmt.Fprintf(gh.out, "Remaining attempts: %d\n", remaining)

	gh.printHint()

	fmt.Fprint(gh.out, "Guessed letters: ")
	for _, letter := range gh.game.GuessedLetters() {
		fmt.Fprintf(gh.out, "%c ", letter)
	}
	fmt.Fprintln(gh.out)
}

func (gh *GameHandler) displayGameResult() {
	if gh.game.IsWon() {
		fmt.Fprintf(gh.out, "🎉 Congratulations! You won! The word was: %s\n", gh.game.Word().Value())
	} else {
		fmt.Fprintf(gh.out, "💀 Game over! The word was: %s\n", gh.game.Word().Value())
	}
}

func (gh *GameHandler) printHint() {
	if gh.showHint {
		fmt.Fprintln(gh.out, "Hint:", gh.game.Word().Hint())
	}
}

func (gh *GameHandler) printMenu() {
	fmt.Fprintln(gh.out, "[1] Start Game")
	fmt.Fprintln(gh.out, "[2] Settings")
	fmt.Fprintln(gh.out, "[3] Exit")
}

func (gh *GameHandler) printError(err error, context string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to %s: %v\n", context, err)
	}
}
