package entities

import (
	"errors"
)

type Level int

const (
	LevelEasy Level = iota + 1
	LevelMedium
	LevelHard
)

var levelNames = map[Level]string{
	LevelEasy:   "Easy",
	LevelMedium: "Medium",
	LevelHard:   "Hard",
}

var levelAttempts = map[Level]int{
	LevelEasy:   10,
	LevelMedium: 8,
	LevelHard:   6,
}

func (l Level) String() string {
	if name, ok := levelNames[l]; ok {
		return name
	}
	return "Unknown"
}

func (l Level) IsValid() bool {
	_, ok := levelNames[l]
	return ok 
}

func (l Level) Attempts() (int, error) {
	if attempts, ok := levelAttempts[l]; ok {
		return attempts, nil
	}

	return 0, errors.New("invalid level")
}

func AllLevels() []Level {
	return []Level{LevelEasy, LevelMedium, LevelHard}
}
