package entities

type GameInfo struct {
	GuessedLetters    []rune
	WordMask          string
	RemainingAttempts int
	Level             string
	Category          string
}

func NewGameInfo(g *Game, cfg *GameConfig) GameInfo {
	return GameInfo{
		GuessedLetters:    g.GuessedLetters(),
		WordMask:          g.WordMask(),
		RemainingAttempts: g.RemainingAttempts(),
		Level:             cfg.Level().String(),
		Category:          cfg.Category().String(),
	}
}
