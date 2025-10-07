package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities/mocks"
)

func TestSelectWord(t *testing.T) {
	mockRepo := mocks.NewMockWordsRepository(t)
	cfg := entities.NewGameConfig(mockRepo)

	cfg.SetLevel(entities.LevelEasy)
	cfg.SetCategory(entities.CategoryAnimals)

	expectedWord := entities.NewWord("dog", "woof")
	mockRepo.EXPECT().
		RandomWord(entities.LevelEasy, entities.CategoryAnimals).
		Return(expectedWord, nil)

	word, err := cfg.SelectWord()
	require.NoError(t, err)
	require.Equal(t, expectedWord, word)
}
