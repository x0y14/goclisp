package interpret

import "fmt"

type RuntimeError struct {
	Kind RuntimeErrorKind
	msg  string
}

func NewRuntimeError(kind RuntimeErrorKind, msg string) *RuntimeError {
	return &RuntimeError{
		Kind: kind,
		msg:  msg,
	}
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("%s: %s", e.Kind.String(), e.msg)
}
