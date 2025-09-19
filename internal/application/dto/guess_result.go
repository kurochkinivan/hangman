package dto

type UserMessage string

const (
	InvalidInput UserMessage = "Invalid Input. P"
	CorrectGuess UserMessage = "Correct!"
	WrongGuess   UserMessage = "Wrong letter."
	LetterUsed   UserMessage = "Letter is already used."
)
