package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/repository/wordslist"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/terminal"
)

func main() {
	wordsRepo := wordslist.NewRepository()

	gameService, err := application.NewGameService(wordsRepo)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch len(os.Args[1:]) {
	case 0:
		h := terminal.NewGameHandler(gameService)

		h.Start()
	case 2:
		word, guessed := os.Args[1], os.Args[2]

		result, err := gameService.SimulateGame(word, guessed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to start interactive mode: %v\n", err)
			os.Exit(1)
		}

		if result.IsWon {
			fmt.Printf("%s;POS\n", result.WordMask)
		} else {
			fmt.Printf("%s;NEG\n", result.WordMask)
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid number of arguments\n")
		os.Exit(1)
	}
}
