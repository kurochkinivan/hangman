package terminal

import (
	"fmt"
	"os"
)

func (gh *GameHandler) setUpGameConfig() error {
	if err := gh.setUpLevel(); err != nil {
		return fmt.Errorf("failed to set up level: %w", err)
	}

	if err := gh.setUpCategory(); err != nil {
		return fmt.Errorf("failed to set up category: %w", err)
	}

	return nil
}

func (gh *GameHandler) setUpLevel() error {
	gh.refreshScreen()

	fmt.Println("Select difficulty level.")

	levels := gh.gameUseCase.Levels()
	for i, l := range levels {
		fmt.Printf("[%d] %s\n", i+1, l.Name)
	}

	var levelID int
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := gh.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpLevel: %v\n", err)
			continue
		}

		if 0 < idx && idx <= len(levels) {
			levelID = levels[idx-1].ID
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}

	return gh.gameUseCase.SetUpLevel(levelID)
}

func (gh *GameHandler) setUpCategory() error {
	gh.refreshScreen()

	fmt.Println("Select words category.")

	categories := gh.gameUseCase.Categories()
	for i, c := range categories {
		fmt.Printf("[%d] %s\n", i+1, c.Name)
	}

	var categoryID int
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := gh.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpCategory: %v\n", err)
			continue
		}

		if 0 < idx && idx <= len(categories) {
			categoryID = categories[idx-1].ID
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}

	return gh.gameUseCase.SetUpCategory(categoryID)
}