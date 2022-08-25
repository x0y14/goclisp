package tokenize

import "fmt"

type SyntaxError struct {
	msg      string
	Subject  string
	Position *Position
}

func NewSyntaxError(position *Position, msg string, subject string) *SyntaxError {
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
	Position       *Position
}

func NewNumberParseError(position *Position, originalString string, err error) *NumberParseError {
	return &NumberParseError{
		OriginalError:  err,
		OriginalString: originalString,
		Position:       position,
	}
}

func (e *NumberParseError) Error() string {
	return e.OriginalError.Error()
}
