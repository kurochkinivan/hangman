package application

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameService struct {
	wordsRepository WordsRepository
}

func NewGameService(wordsRepository WordsRepository) *GameService {
	return &GameService{
		wordsRepository: wordsRepository,
	}
}

type WordsRepository interface {
	RandomWord(config *entities.GameConfig) (*entities.Word, error)
}

func (gs *GameService) CreateGameConfig(level entities.Level, category entities.Category) (*entities.GameConfig, error) {
	config, err := entities.NewGameConfig(level, category)
	if err != nil {
		return nil, fmt.Errorf("failed to create game config: %w", err)
	}

	return config, nil
}

func (gs *GameService) Categories() []entities.Category {
	return entities.AllCategories()
}

func (gs *GameService) Levels() []entities.Level {
	return entities.AllLevels()
}
