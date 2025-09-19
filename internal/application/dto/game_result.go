package dto

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"

type GameResult struct {
	IsWon    bool
	Word     string
	WordMask string
}

func NewGameResult(g *entities.Game) GameResult {
	return GameResult{
		IsWon:    g.Status() == entities.GameStatusWon,
		Word:     g.Word().Value(),
		WordMask: g.WordMask(),
	}
}
