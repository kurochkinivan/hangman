package dto

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"

type Word struct {
	Value string
	Hint  string
}

func NewWord(value, hint string) *Word {
	return &Word{
		Value: value,
		Hint:  hint,
	}
}

func MapDTOToWord(w *Word) *entities.Word {
	return entities.NewWord(w.Value, w.Hint)
}
