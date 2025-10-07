package terminal

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

func (gh *GameHandler) Start() {
	for {
		gh.refreshScreen()
		gh.printMenu()

		choice, _ := gh.readString()

		switch choice {
		case "1":
			if err := gh.startGame(); err != nil {
				gh.printError(err, "start game")
			}

			fmt.Fprintln(gh.out, "Press Enter to return to main menu...")
			gh.reader.ReadString('\n')

		case "2":
			if err := gh.setUpGameConfig(); err != nil {
				gh.printError(err, "set up config")
			}
		case "3":
			return
		}
	}
}

func (gh *GameHandler) startGame() error {
	word, err := gh.config.SelectWord()
	if err != nil {
		return fmt.Errorf("failed to generate word: %w", err)
	}

	gh.game = entities.NewGame(word, gh.config.Level().Attempts())

	for gh.game.InProgress() {
		gh.displayGameState()

		err := gh.play()
		if err != nil {
			gh.printError(err, "failed to play round")
		}
	}

	gh.displayGameState()
	gh.displayGameResult()

	return nil
}

func (gh *GameHandler) play() error {
	fmt.Fprint(gh.out, "Enter a letter (or type 'hint' for a clue): ")

	input, err := gh.readString()
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	if input == "hint" {
		gh.showHint = true
		return nil
	}

	if valid := isValidSingleLetter(input); !valid {
		fmt.Fprintln(gh.out, InvalidInputMsg)
		return nil
	}

	r, _ := utf8.DecodeRuneInString(input)

	if gh.game.IsLetterGuessed(r) {
		fmt.Fprintln(gh.out, "Letter is already used.")
		return nil
	}

	correct := gh.game.GuessLetter(r)
	if correct {
		fmt.Fprintln(gh.out, "Correct!")
	} else {
		fmt.Fprintln(gh.out, "Wrong letter.")
	}

	return nil
}

func isValidSingleLetter(s string) bool {
	if utf8.RuneCountInString(s) != 1 {
		return false
	}

	r, _ := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return false
	}

	if !unicode.IsLetter(r) {
		return false
	}

	return true
}
