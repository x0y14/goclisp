package tokenize

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
)

type SyntaxError struct {
	msg      string
	Subject  string
	Position *data.Position
}

func NewSyntaxError(position *data.Position, msg string, subject string) *SyntaxError {
	return &SyntaxError{
		msg:      msg,
		Subject:  subject,
		Position: position,
	}
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.Subject)
}

type NumberParseError struct {
	OriginalError  error
	OriginalString string
	Position       *data.Position
}

func NewNumberParseError(position *data.Position, originalString string, err error) *NumberParseError {
	return &NumberParseError{
		OriginalError:  err,
		OriginalString: originalString,
		Position:       position,
	}
}

func (e *NumberParseError) Error() string {
	return e.OriginalError.Error()
}
