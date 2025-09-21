package wordslist

import (
	"testing"

	"github.com/stretchr/testify/require"
	ent "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

func TestRandomWord(t *testing.T) {
	tests := []struct {
		name     string
		level    ent.Level
		category ent.Category
		wantErr  bool
		wantWord string
	}{
		{
			name:     "valid easy animals",
			level:    ent.LevelEasy,
			category: ent.CategoryAnimals,
			wantErr:  false,
			wantWord: "cat",
		},
		{
			name:     "invalid category",
			level:    ent.LevelEasy,
			category: ent.CategoryUnknown,
			wantErr:  true,
		},
		{
			name:     "invalid level",
			level:    ent.LevelUnknown,
			category: ent.CategoryAnimals,
			wantErr:  true,
		},
		{
			name:     "empty words list",
			level:    ent.LevelEasy,
			category: ent.CategoryFruitsVegetables,
			wantErr:  true,
		},
	}

	repo := NewRepository()

	animalsEasyBefore := repo.wordsMap[ent.CategoryAnimals][ent.LevelEasy]
	repo.wordsMap[ent.CategoryAnimals][ent.LevelEasy] = []WordEntry{{"cat", ""}}

	fruitsEasyBefore := repo.wordsMap[ent.CategoryFruitsVegetables][ent.LevelEasy]
	repo.wordsMap[ent.CategoryFruitsVegetables][ent.LevelEasy] = []WordEntry{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := ent.NewGameConfig(tt.level, tt.category)
			require.NoError(t, err)

			word, err := repo.RandomWord(cfg)
			if tt.wantErr {
				require.Error(t, err, "should have error")
				return
			}
			require.NoError(t, err)

			require.Equal(t, tt.wantWord, word.Value, "should be equal")
		})
	}

	repo.wordsMap[ent.CategoryAnimals][ent.LevelEasy] = animalsEasyBefore
	repo.wordsMap[ent.CategoryFruitsVegetables][ent.LevelEasy] = fruitsEasyBefore
}
