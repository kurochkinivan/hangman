package terminal

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/terminal/mocks"
)

type PlayTestSuite struct {
	suite.Suite
	handler    *GameHandler
	mockGame   *mocks.MockGame
	mockConfig *mocks.MockConfig
	inputBuf   *bytes.Buffer
	outputBuf  *bytes.Buffer
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(PlayTestSuite))
}

func (s *PlayTestSuite) SetupTest() {
	s.inputBuf = &bytes.Buffer{}
	s.outputBuf = &bytes.Buffer{}
	s.mockGame = mocks.NewMockGame(s.T())
	s.mockConfig = mocks.NewMockConfig(s.T())

	s.handler = &GameHandler{
		game:     s.mockGame,
		config:   s.mockConfig,
		reader:   bufio.NewReader(s.inputBuf),
		out:      s.outputBuf,
		showHint: false,
	}
}

func (s *PlayTestSuite) TestPlay_CorrectGuess() {
	var letter rune = rune('a')
	str := fmt.Sprintf("%c\n", letter)

	_, err := s.inputBuf.WriteString(str)
	s.Require().NoError(err)

	s.mockGame.EXPECT().IsLetterGuessed(letter).Return(false)
	s.mockGame.EXPECT().GuessLetter(letter).Return(true)

	err = s.handler.play()
	s.Require().NoError(err)

	s.Contains(s.outputBuf.String(), "Correct!")
}

func (s *PlayTestSuite) TestPlay_WrongGuess() {
	var letter rune = rune('a')
	str := fmt.Sprintf("%c\n", letter)

	_, err := s.inputBuf.WriteString(str)
	s.Require().NoError(err)

	s.mockGame.EXPECT().IsLetterGuessed(letter).Return(false)
	s.mockGame.EXPECT().GuessLetter(letter).Return(false)

	err = s.handler.play()
	s.Require().NoError(err)

	s.Contains(s.outputBuf.String(), "Wrong letter.")
}

func (s *PlayTestSuite) TestPlay_LetterUsed() {
	var letter rune = rune('a')
	str := fmt.Sprintf("%c\n", letter)

	_, err := s.inputBuf.WriteString(str)
	s.Require().NoError(err)

	s.mockGame.EXPECT().IsLetterGuessed(letter).Return(true)

	err = s.handler.play()
	s.Require().NoError(err)

	s.Contains(s.outputBuf.String(), "Letter is already used.")
}

func TestIsValidSingleLetter(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "valid letter - en",
			s:    "v",
			want: true,
		},
		{
			name: "valid letter - ru",
			s:    "й",
			want: true,
		},
		{
			name: "several letters",
			s:    "ab",
			want: false,
		},
		{
			name: "empty",
			s:    "",
			want: false,
		},
		{
			name: "special symbol",
			s:    "/",
			want: false,
		},
		{
			name: "number",
			s:    "1",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSingleLetter(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
