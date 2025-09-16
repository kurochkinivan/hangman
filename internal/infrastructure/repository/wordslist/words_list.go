package wordslist

import (
	"errors"
	"math/rand/v2"

	ent "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type Repository struct {
	wordsMap map[ent.Category]map[ent.Level][]string
}

// IDEA: Можно использовать [2]int для хранения подсказки
func NewRepository() *Repository {
	return &Repository{
		wordsMap: map[ent.Category]map[ent.Level][]string{
			ent.CategoryAnimals: {
				ent.LevelEasy:   {"cat", "dog", "cow", "pig"},
				ent.LevelMedium: {"lion", "tiger", "zebra", "giraffe"},
				ent.LevelHard:   {"elephant", "rhinoceros", "hippopotamus", "kangaroo"},
			},
			ent.CategoryFruitsVegetables: {
				ent.LevelEasy:   {"apple", "pear", "plum", "carrot"},
				ent.LevelMedium: {"banana", "orange", "pepper", "potato"},
				ent.LevelHard:   {"pomegranate", "asparagus", "artichoke", "cucumber"},
			},
			ent.CategoryCountries: {
				ent.LevelEasy:   {"egypt", "india", "italy", "china"},
				ent.LevelMedium: {"canada", "germany", "brazil", "japan"},
				ent.LevelHard:   {"australia", "kazakhstan", "singapore", "switzerland"},
			},
		},
	}
}

func (r *Repository) RandomWord(config *ent.GameConfig) (*ent.Word, error) {
	categoryWords, ok := r.wordsMap[config.Category()]
	if !ok {
		return nil, errors.New("category not found")
	}

	levelWords, ok := categoryWords[config.Level()]
	if !ok {
		return nil, errors.New("level not found")
	}

	if len(levelWords) == 0 {
		return nil, errors.New("no words are available for this category and level")
	}

	value := levelWords[rand.IntN(len(levelWords))]

	return ent.NewWord(value, config.Level(), config.Category()), nil
}
