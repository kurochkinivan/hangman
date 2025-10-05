package entities

import (
	"math"
	"math/rand/v2"
)

type Level int

const (
	LevelEasy Level = iota + 1
	LevelMedium
	LevelHard
	LevelRandom
	LevelUnknown
)

var playableLevels = []Level{LevelEasy, LevelMedium, LevelHard}

// AllLevels returns the list of levels that the user can choose.
func AllLevels() []Level {
	return []Level{LevelEasy, LevelMedium, LevelHard, LevelRandom}
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
	case LevelRandom:
		return rand.IntN(3) + 5
	case LevelUnknown:
		return math.MaxInt32
	default:
		return 0
	}
}
