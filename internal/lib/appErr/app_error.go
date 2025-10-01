package apperr

import "fmt"

type AppError struct {
	msg string
	op  string
}

func NewAppErr(op string, msg string) AppError {
	return AppError{
		msg: msg,
		op:  op,
	}
}

func (a AppError) Error() string {
	return fmt.Sprintf("%s: %s", a.op, a.msg)
}
