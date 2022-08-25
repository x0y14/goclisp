package parse

import (
	"fmt"
	"github.com/x0y14/goclisp/tokenize"
)

type SyntaxError struct {
	Token *tokenize.Token
	msg   string
}

func NewSyntaxError(msg string, tok *tokenize.Token) *SyntaxError {
	return &SyntaxError{
		Token: tok,
		msg:   msg,
	}
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("[%d:%d] %s: %s",
		e.Token.Position.LineNo,
		e.Token.Position.LpBegin,
		e.msg, e.Token.Str)
}
