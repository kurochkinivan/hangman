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

var levelNames = map[Level]string{
	LevelEasy:    "Easy",
	LevelMedium:  "Medium",
	LevelHard:    "Hard",
	LevelRandom:  "Random",
}

func (l Level) IsValid() bool {
	_, ok := levelNames[l]
	return ok
}

func (l Level) String() string {
	if name, ok := levelNames[l]; ok {
		return name
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

func AllLevels() []Level {
	levels := make([]Level, 0, len(levelNames))

	for level := range levelNames {
		levels = append(levels, level)
	}

	return levels
}

func RandomLevel() Level {
	levels := []Level{LevelEasy, LevelMedium, LevelHard}
	return levels[rand.IntN(len(levels))]
}
