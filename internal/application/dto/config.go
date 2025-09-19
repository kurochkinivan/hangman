package dto

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type Level struct {
	ID   int
	Name string
}

func MapLevelToDTO(l entities.Level) Level {
	return Level{
		ID:   int(l),
		Name: l.String(),
	}
}

type Category struct {
	ID   int
	Name string
}

func MapCategoryToDTO(c entities.Category) Category {
	return Category{
		ID:   int(c),
		Name: c.String(),
	}
}
