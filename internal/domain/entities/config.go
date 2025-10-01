package entities

import (
	apperr "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/lib/appErr"
)

type GameConfig struct {
	level       Level
	category    Category
	maxAttempts int
}

func NewGameConfig(level Level, category Category) (*GameConfig, error) {
	if !level.IsValid() {
		return nil, apperr.NewAppErr("NewGameConfig", "invalid level")
	}

	if !category.IsValid() {
		return nil, apperr.NewAppErr("NewGameConfig", "invalid category")
	}

	return &GameConfig{
		level:       level,
		category:    category,
		maxAttempts: level.Attempts(),
	}, nil
}

func (gc *GameConfig) Category() Category {
	return gc.category
}

func (gc *GameConfig) SetCategory(category Category) error {
	if !category.IsValid() {
		return apperr.NewAppErr("NewGameConfig", "invalid category")
	}

	gc.category = category

	return nil
}

func (gc *GameConfig) Level() Level {
	return gc.level
}

func (gc *GameConfig) SetLevel(level Level) error {
	if !level.IsValid() {
		return apperr.NewAppErr("NewGameConfig", "invalid level")
	}

	gc.level = level
	gc.maxAttempts = level.Attempts()

	return nil
}

func (gc *GameConfig) MaxAttempts() int {
	return gc.maxAttempts
}
