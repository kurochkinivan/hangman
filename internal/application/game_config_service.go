package application

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)


func (gs *GameService) SetUpLevel(levelID int) error {
	level := entities.Level(levelID)
	if !level.IsValid() {
		return fmt.Errorf("invalid level")
	}

	return gs.globalConfig.SetLevel(level)
}

func (gs *GameService) SetUpCategory(categoryID int) error {
	category := entities.Category(categoryID)
	if !category.IsValid() {
		return fmt.Errorf("invalid category")
	}

	return gs.globalConfig.SetCategory(category)
}

func (gs *GameService) Categories() []dto.Category {
	categories := entities.AllCategories()
	result := make([]dto.Category, 0, len(categories))

	for _, c := range categories {
		result = append(result, dto.MapCategoryToDTO(c))
	}

	return result
}

func (gs *GameService) Levels() []dto.Level {
	levels := entities.AllLevels()
	result := make([]dto.Level, 0, len(levels))

	for _, l := range levels {
		result = append(result, dto.MapLevelToDTO(l))
	}

	return result
}
