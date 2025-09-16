package terminal

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

func (g *GameHandler) setUpGameConfig() (*entities.GameConfig, error) {
	level := g.setUpLevel()
	category := g.setUpCategory()

	config, err := g.gameConfigUseCase.CreateGameConfig(level, category)
	if err != nil {
		return nil, fmt.Errorf("setUpGameConfig: failed to create game config: %w", err)
	}

	return config, nil
}

// TODO: Здесь можно выбирать рандомный параметр, а дальше уже решать
func (g *GameHandler) setUpCategory() entities.Category {
	g.prepareScreen()

	fmt.Println("Select words category.")

	categories := g.gameConfigUseCase.Categories()
	for i, c := range categories {
		fmt.Printf("[%d] %s\n", i+1, c)
	}

	var category entities.Category
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := g.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpCategory: %v\n", err)
		}

		if 0 < idx && idx <= len(categories) {
			category = categories[idx-1]
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}

	return category
}

// TODO: Здесь можно выбирать рандомный параметр, а дальше уже решать
func (g *GameHandler) setUpLevel() entities.Level {
	g.prepareScreen()

	fmt.Println("Select difficulty level.")

	levels := g.gameConfigUseCase.Levels()
	for i, l := range levels {
		fmt.Printf("[%d] %s\n", i+1, l)
	}

	var level entities.Level
	for {
		fmt.Println(EnterYourChoiceMsg)

		idx, err := g.readInt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "setUpLevel: %v\n", err)
		}

		if 0 < idx && idx <= len(levels) {
			level = levels[idx-1]
			break
		} else {
			fmt.Println(InvalidInputMsg)
		}
	}

	return level
}
