package simulate

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

func TestSimulateGame(t *testing.T) {
	tests := []struct {
		name          string
		word, guessed string
		want          entities.GameResult
		wantErr       bool
	}{
		{
			name:    "game win - en",
			word:    "cat",
			guessed: "cat",
			want:    entities.NewGameResult("cat", true),
			wantErr: false,
		},
		{
			name:    "game win - ru",
			word:    "анатолий",
			guessed: "анатолий",
			want:    entities.NewGameResult("анатолий", true),
			wantErr: false,
		},
		{
			name:    "game loose - en",
			word:    "cat",
			guessed: "caa",
			want:    entities.NewGameResult("ca*", false),
			wantErr: false,
		},
		{
			name:    "game loose - ru",
			word:    "армагедон",
			guessed: "армагееее",
			want:    entities.NewGameResult("армаге***", false),
			wantErr: false,
		},
		{
			name:    "different languages",
			word:    "cat",
			guessed: "дог",
			want:    entities.NewGameResult("***", false),
			wantErr: false,
		},
		{
			name:    "different length",
			word:    "cat",
			guessed: "catcat",
			want:    entities.NewGameResult("***", false),
			wantErr: true,
		},
		{
			name:    "empty word",
			word:    "",
			guessed: "",
			want:    entities.NewGameResult("", true),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gh := NewGameHandler(tt.word, tt.guessed)
			result, err := gh.simulateGame()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, result)
		})
	}
}
