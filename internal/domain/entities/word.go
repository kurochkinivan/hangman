package entities

import "strings"

type Word struct {
	value    string
	level    Level
	category Category
}

func NewWord(value string, level Level, category Category) *Word {
	return &Word{
		value:    strings.ToLower(value),
		level:    level,
		category: category,
	}
}

func (w *Word) Value() string {
	return w.value
}

func (w *Word) Contains(r rune) bool {
	return strings.ContainsRune(w.value, r)
}

