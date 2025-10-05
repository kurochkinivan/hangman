package wordslist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ent "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/repository/wordslist/mocks"
)

func TestRandomWord_HappyPath(t *testing.T) {
	expectedWord := Word{Value: "cat", Hint: "meow"}
	wl := WordsList{
		Animals: CategoryData{
			Easy: []Word{expectedWord, {Value: "dog", Hint: "woof"}},
		},
	}

	mockRandSel := mocks.NewMockRandomSelector[Word](t)
	mockRandSel.EXPECT().Choose(wl.Animals.Easy).Return(expectedWord)

	repo := NewRepository(wl, mockRandSel)

	wordDTO, err := repo.RandomWord(ent.LevelEasy, ent.CategoryAnimals)
	require.NoError(t, err)

	assert.Equal(t, expectedWord.Value, wordDTO.Value())
	assert.Equal(t, expectedWord.Hint, wordDTO.Hint())
}

func TestRandomWord_EmptyCategory(t *testing.T) {
	wl := WordsList{}

	mockRandSel := mocks.NewMockRandomSelector[Word](t)
	repo := NewRepository(wl, mockRandSel)

	w, err := repo.RandomWord(ent.LevelEasy, ent.CategoryAnimals)

	assert.Error(t, err)
	assert.Empty(t, w)
}

func TestRandomWord_EmptyLevel(t *testing.T) {
	wl := WordsList{
		Animals: CategoryData{},
	}

	mockRandSel := mocks.NewMockRandomSelector[Word](t)
	repo := NewRepository(wl, mockRandSel)

	w, err := repo.RandomWord(ent.LevelEasy, ent.CategoryAnimals)

	assert.Error(t, err)
	assert.Empty(t, w)
}

func TestRandomWord_NoWords(t *testing.T) {
	wl := WordsList{
		Animals: CategoryData{
			Easy: []Word{},
		},
	}

	mockRandSel := mocks.NewMockRandomSelector[Word](t)
	repo := NewRepository(wl, mockRandSel)

	w, err := repo.RandomWord(ent.LevelEasy, ent.CategoryAnimals)

	assert.Error(t, err)
	assert.Empty(t, w)
}
