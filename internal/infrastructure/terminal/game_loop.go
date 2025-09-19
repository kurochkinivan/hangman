package terminal

import (
	"fmt"
	"os"
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
	err := gh.gameUseCase.StartNewGame()
	if err != nil {
		return fmt.Errorf("failed to start new game: %w", err)
	}

	for gh.gameUseCase.InProgress() {
		err := gh.play()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			fmt.Println(InvalidInputMsg)
		}
	}

	gh.displayGameState(gh.gameUseCase.GameInfo())
	gh.displayGameResult(gh.gameUseCase.GameResult())

	return nil
}

func (gh *GameHandler) play() error {
	gh.displayGameState(gh.gameUseCase.GameInfo())

	fmt.Println("Enter a letter (or type 'hint' for a clue):")
	input, err := gh.readString()
	if err != nil {
		return fmt.Errorf("failed to read string: %v", err)
	}

	if input == "hint" {
		gh.showHint = true
		return nil
	}

	message, err := gh.gameUseCase.GuessLetter(input)
	if err != nil {
		return fmt.Errorf("failed to guess letter: %w", err)
	}
	fmt.Println("Message:", message)

	return nil
}
