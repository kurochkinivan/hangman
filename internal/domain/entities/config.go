package entities

import (
	"errors"
)

type GameConfig struct {
	level    Level
	category Category
}

func NewGameConfig(level Level, category Category) (*GameConfig, error) {
	if !level.IsValid() {
		return nil, errors.New("invalid level")
	}

	if !category.IsValid() {
		return nil, errors.New("invalid category")
	}

	return &GameConfig{
		level:    level,
		category: category,
	}, nil
}

func (gc *GameConfig) Category() Category {
	return gc.category
}

func (gc *GameConfig) Level() Level {
	return gc.level
}