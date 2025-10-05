package entities

import "strings"

type Word struct {
	value string
	hint  string
}

func NewWord(value string, hint string) Word {
	return Word{
		value: strings.ToLower(value),
		hint:  hint,
	}
}

func (w Word) Value() string {
	return w.value
}

func (w Word) Hint() string {
	return w.hint
}

func (w Word) Contains(r rune) bool {
	return strings.ContainsRune(w.value, r)
}
