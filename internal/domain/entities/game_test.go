package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Нужен ли этот тест? 
func TestNewGameValidation(t *testing.T) {
	tests := []struct {
		name    string
		word    *Word
		config  *GameConfig
		wantErr bool
	}{
		{
			name:    "nil word",
			word:    nil,
			config:  &GameConfig{},
			wantErr: true,
		},
		{
			name:    "nil config",
			word:    &Word{},
			config:  nil,
			wantErr: true,
		},
		{
			name:    "valid",
			word:    &Word{},
			config:  &GameConfig{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGame(tt.word, tt.config)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGameFlow(t *testing.T) {
	tests := []struct {
		name                  string
		word                  string
		maxAttempts           int
		guesses               []rune
		wantMask              string
		wantRemainingAttempts int
		wantStatus            GameStatus
		wantGuessedRune       []rune
	}{
		{
			name:                  "Loose en",
			word:                  "go",
			maxAttempts:           2,
			guesses:               []rune{'x', 'w'},
			wantMask:              "**",
			wantRemainingAttempts: 0,
			wantStatus:            GameStatusLost,
			wantGuessedRune:       []rune{'w', 'x'},
		},
		{
			name:                  "Win en",
			word:                  "go",
			maxAttempts:           2,
			guesses:               []rune{'o', 'g'},
			wantMask:              "go",
			wantRemainingAttempts: 2,
			wantStatus:            GameStatusWon,
			wantGuessedRune:       []rune{'g', 'o'},
		},
		{
			name:                  "In progress en",
			word:                  "go",
			maxAttempts:           4,
			guesses:               []rune{'w', 'x', 'g'},
			wantMask:              "g*",
			wantRemainingAttempts: 2,
			wantStatus:            GameStatusInProgress,
			wantGuessedRune:       []rune{'g', 'w', 'x'},
		},
		{
			name:                  "Loose ru",
			word:                  "МАМА",
			maxAttempts:           2,
			guesses:               []rune{'ф', 'в'},
			wantMask:              "****",
			wantRemainingAttempts: 0,
			wantStatus:            GameStatusLost,
			wantGuessedRune:       []rune{'в', 'ф'},
		},
		{
			name:                  "Win ru",
			word:                  "МАМА",
			maxAttempts:           2,
			guesses:               []rune{'м', 'а'},
			wantMask:              "мама",
			wantRemainingAttempts: 2,
			wantStatus:            GameStatusWon,
			wantGuessedRune:       []rune{'а', 'м'},
		},
		{
			name:                  "In progress ru",
			word:                  "МАМА",
			maxAttempts:           4,
			guesses:               []rune{'а', 'б', 'в'},
			wantMask:              "*а*а",
			wantRemainingAttempts: 2,
			wantStatus:            GameStatusInProgress,
			wantGuessedRune:       []rune{'а', 'б', 'в'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWord(tt.word, "")

			g, err := NewGame(w, &GameConfig{
				maxAttempts: tt.maxAttempts,
			})
			require.NoError(t, err)

			for _, guess := range tt.guesses {
				g.GuessLetter(guess)
			}

			assert.Equal(t, tt.wantMask, g.WordMask())

			assert.Equal(t, tt.wantStatus, g.Status())

			assert.Equal(t, tt.wantGuessedRune, g.GuessedLetters())

			assert.Equal(t, tt.wantRemainingAttempts, g.RemainingAttempts())
		})
	}
}
