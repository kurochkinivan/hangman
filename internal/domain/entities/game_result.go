package entities

type GameResult struct {
	wordMask string
	isWon    bool
}

func NewGameResult(WordMask string, IsWon bool) GameResult {
	return GameResult{
		wordMask: WordMask,
		isWon:    IsWon,
	}
}

func (gr GameResult) WordMask() string {
	return gr.wordMask
}

func (gr GameResult) IsWon() bool {
	return gr.isWon
}
