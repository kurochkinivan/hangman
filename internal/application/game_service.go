package application

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application/dto"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameService struct {
	game            *entities.Game
	globalConfig    *entities.GameConfig // Global config, can be set to 'random'
	wordsRepository WordsRepository
}

func NewGameService(wordsRepository WordsRepository) (*GameService, error) {
	globalConfig, err := entities.NewGameConfig(entities.LevelRandom, entities.CategoryRandom)
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}

	return &GameService{
		globalConfig:    globalConfig,
		wordsRepository: wordsRepository,
	}, nil
}

type WordsRepository interface {
	RandomWord(config *entities.GameConfig) (*dto.Word, error)
}

func (gs *GameService) StartNewGame() error {
	level := gs.globalConfig.Level()
	if level == entities.LevelRandom {
		level = entities.RandomLevel()
	}

	category := gs.globalConfig.Category()
	if category == entities.CategoryRandom {
		category = entities.RandomCategory()
	}

	gameConfig, err := entities.NewGameConfig(level, category)
	if err != nil {
		return fmt.Errorf("failed to create game config")
	}

	word, err := gs.wordsRepository.RandomWord(gameConfig)
	if err != nil {
		return fmt.Errorf("failed to craete random word: %w", err)
	}

	return gs.loadGame(word, gameConfig)
}

func (gs *GameService) SimulateGame(word string, guessed string) (dto.GameResult, error) {
	if utf8.RuneCountInString(word) != utf8.RuneCountInString(guessed) {
		return dto.GameResult{}, errors.New("lengths of given word and guessed word do not match")
	}

	cfg, err := entities.NewGameConfig(entities.LevelUnknown, entities.CategoryUnknown)
	if err != nil {
		return dto.GameResult{}, fmt.Errorf("failed to create new game config: %w", err)
	}

	err = gs.loadGame(dto.NewWord(word, ""), cfg)
	if err != nil {
		return dto.GameResult{}, fmt.Errorf("failed to start game: %w", err)
	}

	for _, r := range guessed {
		gs.GuessLetter(string(r))
	}

	return gs.GameResult(), nil
}

func (gs *GameService) loadGame(word *dto.Word, gameConfig *entities.GameConfig) error {
	var err error

	gs.game, err = entities.NewGame(dto.MapDTOToWord(word), gameConfig)
	if err != nil {
		return fmt.Errorf("failed to craete new game: %w", err)
	}

	return nil
}

func (gs *GameService) GuessLetter(input string) (dto.UserMessage, error) {
	if utf8.RuneCountInString(input) != 1 {
		return dto.InvalidInput, errors.New("rune count in input has to be 1")
	}

	r, size := utf8.DecodeRuneInString(input)
	if r == utf8.RuneError {
		if size == 1 {
			return dto.InvalidInput, fmt.Errorf("invalid encoding, failed to decode %q", input)
		}
		return dto.InvalidInput, errors.New("input is empty")
	}

	if !unicode.IsLetter(r) {
		return dto.InvalidInput, errors.New("rune has to be letter")
	}

	if gs.game.IsLetterGuessed(r) {
		return dto.LetterUsed, nil
	}

	correct := gs.game.GuessLetter(r)
	if correct {
		return dto.CorrectGuess, nil
	}

	return dto.WrongGuess, nil
}

func (gs *GameService) Hint() string {
	return gs.game.Word().Hint()
}

func (gs *GameService) GameInfo() dto.GameInfo {
	return dto.NewGameInfo(gs.game, gs.game.Config())
}

func (gs *GameService) GameResult() dto.GameResult {
	return dto.NewGameResult(gs.game)
}

func (gs *GameService) InProgress() bool {
	return gs.game.Status() == entities.GameStatusInProgress
}
