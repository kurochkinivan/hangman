package wordslist

import (
	"errors"
	"math/rand/v2"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
	ent "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type WordEntry struct {
	Value string
	Hint  string
}

type Repository struct {
	wordsMap map[ent.Category]map[ent.Level][]WordEntry
}

func NewRepository() *Repository {
	return &Repository{
		wordsMap: map[ent.Category]map[ent.Level][]WordEntry{
			ent.CategoryAnimals: {
				ent.LevelEasy: {
					{"cat", "domestic pet"},
					{"dog", "man's best friend"},
					{"cow", "gives milk"},
					{"pig", "lives on a farm"},
				},
				ent.LevelMedium: {
					{"lion", "king of the jungle"},
					{"tiger", "striped predator"},
					{"zebra", "striped horse"},
					{"giraffe", "long neck"},
				},
				ent.LevelHard: {
					{"elephant", "largest land animal"},
					{"rhinoceros", "has a horn"},
					{"hippopotamus", "lives near rivers"},
					{"kangaroo", "jumps and has a pouch"},
				},
			},
			ent.CategoryFruitsVegetables: {
				ent.LevelEasy: {
					{"apple", "keeps the doctor away"},
					{"pear", "bell-shaped fruit"},
					{"plum", "small purple fruit"},
					{"carrot", "orange vegetable"},
				},
				ent.LevelMedium: {
					{"banana", "yellow curved fruit"},
					{"orange", "citrus fruit"},
					{"pepper", "can be sweet or spicy"},
					{"potato", "grows underground"},
				},
				ent.LevelHard: {
					{"pomegranate", "many red seeds inside"},
					{"asparagus", "green spring vegetable"},
					{"artichoke", "edible flower bud"},
					{"cucumber", "green salad vegetable"},
				},
			},
			ent.CategoryCountries: {
				ent.LevelEasy: {
					{"egypt", "pyramids"},
					{"india", "Taj Mahal"},
					{"italy", "pizza & pasta"},
					{"china", "Great Wall"},
				},
				ent.LevelMedium: {
					{"canada", "maple leaf"},
					{"germany", "beer & sausages"},
					{"brazil", "Rio carnival"},
					{"japan", "samurai & sushi"},
				},
				ent.LevelHard: {
					{"australia", "kangaroos and koalas"},
					{"kazakhstan", "largest landlocked country"},
					{"singapore", "city-state in Asia"},
					{"switzerland", "Alps and chocolate"},
				},
			},
		},
	}
}

func (r *Repository) RandomWord(config *ent.GameConfig) (*dto.Word, error) {
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

	wordEntry := levelWords[rand.IntN(len(levelWords))]

	return dto.NewWord(wordEntry.Value, wordEntry.Hint), nil
}
