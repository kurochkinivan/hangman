package entities

import (
	"errors"
)

// TODO: Подумать про рандом.
type Level string

const (
	LevelEasy   Level = "Easy"
	LevelMedium Level = "Medium"
	LevelHard   Level = "Hard"
)

var (
	levels   = []Level{LevelEasy, LevelMedium, LevelHard}
	attempts = map[Level]int{
		LevelEasy:   10,
		LevelMedium: 8,
		LevelHard:   6,
	}
)

func (l Level) Attempts() (int, error) {
	if n, ok := attempts[l]; ok {
		return n, nil
	}
	return -1, errors.New("unknown level")
}

func AllLevels() []Level {
	return append([]Level(nil), levels...)
}
