package dto

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"

type GameInfo struct {
	GuessedLetters    []rune
	WordMask          string
	RemainingAttempts int
	Level             string
	Category          string
}

func NewGameInfo(g *entities.Game, cfg *entities.GameConfig) GameInfo {
	return GameInfo{
		GuessedLetters:    g.GuessedLetters(),
		WordMask:          g.WordMask(),
		RemainingAttempts: g.RemainingAttempts(),
		Level:             cfg.Level().String(),
		Category:          cfg.Category().String(),
	}
}
