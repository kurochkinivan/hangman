package terminal

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
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

	levels := entities.AllLevels()
	for i, l := range levels {
		fmt.Printf("[%d] %s\n", i+1, l.String())
	}

	var level entities.Level
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := gh.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpLevel: %v\n", err)
			continue
		}

		if 0 < idx && idx <= len(levels) {
			level = levels[idx-1]
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}
	gh.config.SetLevel(level)

	return nil
}

func (gh *GameHandler) setUpCategory() error {
	gh.refreshScreen()

	fmt.Println("Select words category.")

	categories := entities.AllCategories()
	for i, cat := range categories {
		fmt.Printf("[%d] %s\n", i+1, cat.String())
	}

	var category entities.Category
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := gh.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpCategory: %v\n", err)
			continue
		}

		if 0 < idx && idx <= len(categories) {
			category = categories[idx-1]
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}
	gh.config.SetCategory(category)

	return nil
}
