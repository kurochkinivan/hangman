package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/repository/wordslist"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/simulate"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/terminal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/pkg/random"
)

func main() {
	wl, err := wordslist.LoadWordsListFromYAML()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load repository from yaml: %v\n", err.Error())
		os.Exit(1)
	}
	randSelector := random.New[wordslist.Word]()
	wordsRepo := wordslist.NewRepository(wl, randSelector)

	switch len(os.Args[1:]) {
	case 0:
		gh := terminal.NewGameHandler(wordsRepo, os.Stdin, os.Stdout)
		gh.Start()
	case 2:
		word, guessed := os.Args[1], os.Args[2]

		gh := simulate.NewGameHandler(word, guessed)
		gh.Start()
	default:
		fmt.Fprintf(os.Stderr, "Invalid number of arguments\n")
		os.Exit(1)
	}
}
