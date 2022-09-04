package server

import (
	"errors"
	"fmt"
)


// エラー発生のオペレーションを特定する様

// errors.Asで
func HandleSignup() error {
	if err := createUser(); err != nil {
		var e *Error
		if errors.As(err, e) {
			return e
		}

		return err
	}

	return nil
}

func createUser() error {
	return errors.New("")
}

type Error struct {
	op string
	err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("op: %s, err: %s", e.op, e.err.Error())
}

func (e *Error) Unwrap() error {
	return e.err
}

