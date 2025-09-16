package main

import (
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/repository/wordslist"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/terminal"
)

func main() {
	f, err := os.OpenFile("stderr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		os.Stderr = f
	}

	repo := wordslist.NewRepository()

	service := application.NewGameService(repo)

	h := terminal.NewGameHandler(service)
	
	h.Start()
}
