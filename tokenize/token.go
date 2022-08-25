package tokenize

type Token struct {
	Kind     TokenKind
	Position *Position

	Num float64
	Str string

	Next *Token
}

func NewToken(kind TokenKind, position *Position, num float64, str string) *Token {
	return &Token{
		Kind:     kind,
		Position: position,
		Num:      num,
		Str:      str,
		Next:     nil,
	}
}

func NewTokenEof(cur *Token, position *Position) *Token {
	tok := NewToken(Eof, position, 0, "")
	cur.Next = tok
	return tok
}

func NewTokenReserved(cur *Token, position *Position, str string) *Token {
	tok := NewToken(Reserved, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenIdent(cur *Token, position *Position, str string) *Token {
	tok := NewToken(Ident, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenString(cur *Token, position *Position, str string) *Token {
	tok := NewToken(String, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenNumber(cur *Token, position *Position, num float64, str string) *Token {
	tok := NewToken(Number, position, num, str)
	cur.Next = tok
	return tok
}
