package wordslist

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
	ent "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/lib/errs"
)

type Repository struct {
	wl      WordsList
	randSel RandomSelector[Word]
}

type RandomSelector[T any] interface {
	Choose(items []T) T
}

func NewRepository(wl WordsList, randSel RandomSelector[Word]) *Repository {
	return &Repository{
		wl:      wl,
		randSel: randSel,
	}
}

func (r *Repository) RandomWord(level ent.Level, category ent.Category) (entities.Word, error) {
	categoryData, err := r.wl.categoryDataFromCategory(category)
	if err != nil {
		return entities.Word{}, fmt.Errorf("failed to get category data from category: %w", err)
	}

	levelWords, err := categoryData.wordEntriesFromLevel(level)
	if err != nil {
		return entities.Word{}, fmt.Errorf("failed to get words from level: %w", err)
	}

	if len(levelWords) == 0 {
		return entities.Word{}, errs.NewAppErr("Repository.RandomWord", "no words are available for this category and level")
	}

	wordEntry := r.randSel.Choose(levelWords)

	return entities.NewWord(wordEntry.Value, wordEntry.Hint), nil
}

func (wl WordsList) categoryDataFromCategory(category ent.Category) (CategoryData, error) {
	switch category {
	case ent.CategoryAnimals:
		return wl.Animals, nil
	case ent.CategoryFruitsVegetables:
		return wl.FruitsVegetables, nil
	case ent.CategoryCountries:
		return wl.Countries, nil
	default:
		return CategoryData{}, fmt.Errorf("unknown category %s", category)
	}
}

func (cat CategoryData) wordEntriesFromLevel(level ent.Level) ([]Word, error) {
	switch level {
	case ent.LevelEasy:
		return cat.Easy, nil
	case ent.LevelMedium:
		return cat.Medium, nil
	case ent.LevelHard:
		return cat.Hard, nil
	default:
		return nil, fmt.Errorf("unknown level %s", level)
	}
}
