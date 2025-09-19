package entities

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Level int

const (
	LevelEasy Level = iota + 1
	LevelMedium
	LevelHard
	LevelRandom
	LevelUnknown
)

var (
	allLevels        = []Level{LevelEasy, LevelMedium, LevelHard, LevelRandom, LevelUnknown} // все существующие
	selectableLevels = []Level{LevelEasy, LevelMedium, LevelHard, LevelRandom}               // то, что можно выбрать
	playableLevels   = []Level{LevelEasy, LevelMedium, LevelHard}
)

func (l Level) IsValid() bool {
	return slices.Contains(allLevels, l)
}

func AllLevels() []Level {
	return selectableLevels
}

func RandomLevel() Level {
	return playableLevels[rand.IntN(len(playableLevels))]
}

func (l Level) String() string {
	switch l {
	case LevelEasy:
		return "Easy"
	case LevelMedium:
		return "Medium"
	case LevelHard:
		return "Hard"
	case LevelRandom:
		return "Random"
	}
	return "Invalid level"
}

func (l Level) Attempts() int {
	switch l {
	case LevelEasy:
		return 7
	case LevelMedium:
		return 6
	case LevelHard:
		return 5
	case LevelUnknown:
		return math.MaxInt32
	default:
		return 0
	}
}
