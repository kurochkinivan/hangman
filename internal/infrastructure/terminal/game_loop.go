package terminal

import (
	"fmt"
	"os"
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

			fmt.Println("Press Enter to return to main menu...")
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
	word, err := gh.config.GenerateWord()
	if err != nil {
		return fmt.Errorf("failed to generate word: %w", err)
	}

	gh.game = entities.NewGame(word, gh.config.Level().Attempts())

	for gh.game.InProgress() {
		err := gh.play()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	gh.displayGameState()
	gh.displayGameResult()

	return nil
}

func (gh *GameHandler) play() error {
	gh.displayGameState()
	fmt.Print("Enter a letter (or type 'hint' for a clue): ")

	input, err := gh.readString()
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	if input == "hint" {
		gh.showHint = true
		return nil
	}

	if valid, err := isValidSingleLetter(input); !valid {
		fmt.Println(InvalidInputMsg)
		return err
	}

	r, _ := utf8.DecodeRuneInString(input)

	if gh.game.IsLetterGuessed(r) {
		fmt.Println("Letter is already used.")
		return nil
	}

	correct := gh.game.GuessLetter(r)
	if correct {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong letter.")
	}

	return nil
}

func isValidSingleLetter(s string) (bool, error) {
	if utf8.RuneCountInString(s) != 1 {
		return false, nil
	}

	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		if size == 1 {
			return false, fmt.Errorf("invalid encoding, failed to decode %q", s)
		}
		return false, nil
	}

	if !unicode.IsLetter(r) {
		return false, nil
	}

	return true, nil
}
