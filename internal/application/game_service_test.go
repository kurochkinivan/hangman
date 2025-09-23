package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

func TestGuessLetter(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		word    string
		wantMsg dto.UserMessage
		wantErr bool
	}{
		{
			name:    "valid input: lowercase",
			input:   "м",
			word:    "мама",
			wantMsg: dto.CorrectGuess,
			wantErr: false,
		},
		{
			name:    "valid input: uppercase",
			input:   "М",
			word:    "мама",
			wantMsg: dto.CorrectGuess,
			wantErr: false,
		},
		{
			name:    "invalid input: wrong letter",
			input:   "ё",
			word:    "мама",
			wantMsg: dto.WrongGuess,
			wantErr: false,
		},
		{
			name:    "invalid input: number",
			input:   "1",
			word:    "мама",
			wantMsg: dto.InvalidInput,
			wantErr: true,
		},
		{
			name:    "invalid input: special digit",
			input:   "/",
			word:    "мама",
			wantMsg: dto.InvalidInput,
			wantErr: true,
		},
		{
			name:    "invalid input: empty",
			input:   "",
			word:    "мама",
			wantMsg: dto.InvalidInput,
			wantErr: true,
		},
		{
			name:    "invalid input: multi letter",
			input:   "fb",
			word:    "мама",
			wantMsg: dto.InvalidInput,
			wantErr: true,
		},
		{
			name:    "invalid input: special symbol",
			input:   "\n",
			word:    "мама",
			wantMsg: dto.InvalidInput,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := prepareGameService(t, tt.word, entities.LevelUnknown, entities.CategoryUnknown)

			msg, err := gs.GuessLetter(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantMsg, msg)
		})
	}
}

func TestGuessLetter_LetterUsed(t *testing.T) {
	word := "МАМА"
	gs := prepareGameService(t, word, entities.LevelUnknown, entities.CategoryUnknown)

	letter := "м"

	msg, err := gs.GuessLetter(letter)
	require.NoError(t, err)
	assert.Equal(t, dto.CorrectGuess, msg)

	msg, err = gs.GuessLetter(letter)
	require.NoError(t, err)
	assert.Equal(t, dto.LetterUsed, msg)
}

func TestGuessLetter_AttemptsChange(t *testing.T) {
	word := "подкрадули"
	level := entities.LevelEasy

	attempts := level.Attempts()
	gs := prepareGameService(t, word, level, entities.CategoryUnknown)

	correctLetter := "п"
	gs.GuessLetter(correctLetter)
	require.Equal(t, attempts, gs.game.RemainingAttempts())

	usedLetter := correctLetter
	gs.GuessLetter(usedLetter)
	require.Equal(t, attempts, gs.game.RemainingAttempts())

	severalLetters := "по"
	gs.GuessLetter(severalLetters)
	require.Equal(t, attempts, gs.game.RemainingAttempts())

	wrongLetter := "я"
	gs.GuessLetter(wrongLetter)
	attempts--
	require.Equal(t, attempts, gs.game.RemainingAttempts())

	noLetters := ""
	gs.GuessLetter(noLetters)
	require.Equal(t, attempts, gs.game.RemainingAttempts())
}

func TestSimulateGame(t *testing.T) {
	tests := []struct {
		name         string
		word         string
		guessed      string
		wantErr      bool
		wantIsWon    bool
		wantWordMask string
	}{
		{
			name:         "equal words",
			word:         "волокно",
			guessed:      "волокно",
			wantErr:      false,
			wantIsWon:    true,
			wantWordMask: "волокно",
		},
		{
			name:         "not partially equal words",
			word:         "волокно",
			guessed:      "толокно",
			wantErr:      false,
			wantIsWon:    false,
			wantWordMask: "*олокно",
		},
		{
			name:         "not equal words",
			word:         "аааа",
			guessed:      "оооо",
			wantErr:      false,
			wantIsWon:    false,
			wantWordMask: "****",
		},
		{
			name:         "different words length",
			word:         "волокно",
			guessed:      "БОЛЬНО В НОГЕ АААААААААААААА",
			wantErr:      true,
			wantIsWon:    false,
			wantWordMask: "",
		},
		{
			name:         "different length",
			word:         "волокно",
			guessed:      "БОЛЬНО В НОГЕ АААААААААААААА",
			wantErr:      true,
			wantIsWon:    false,
			wantWordMask: "",
		},
		{
			name:         "different order",
			word:         "окок",
			guessed:      "коко",
			wantErr:      false,
			wantIsWon:    true,
			wantWordMask: "окок",
		},
		{
			name:         "different languages",
			word:         "мама",
			guessed:      "mama",
			wantErr:      false,
			wantIsWon:    false,
			wantWordMask: "****",
		},
	}

	gs := &GameService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := gs.SimulateGame(tt.word, tt.guessed)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tt.wantIsWon, result.IsWon)
			assert.Equal(t, tt.wantWordMask, result.WordMask)
		})
	}
}

func prepareGameService(t *testing.T, word string, level entities.Level, category entities.Category) *GameService {
	t.Helper()
	t.Parallel()

	gs := &GameService{}

	cfg, err := entities.NewGameConfig(level, category)
	require.NoError(t, err)

	err = gs.loadGame(dto.NewWord(word, ""), cfg)
	require.NoError(t, err)

	return gs
}
