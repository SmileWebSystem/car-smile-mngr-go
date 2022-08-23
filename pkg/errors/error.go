package errors

import "fmt"

type Error struct {
	code    string
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %v", e.code, e.message)
}

//func (e Error) getCode() string {
//	return e.code
//}

func NewError(code string, message string) error {
	return Error{code: code, message: message}
}
