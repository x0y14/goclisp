package data

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
	tok := NewToken(TkEof, position, 0, "")
	cur.Next = tok
	return tok
}

func NewTokenReserved(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkReserved, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenIdent(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkIdent, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenString(cur *Token, position *Position, str string) *Token {
	tok := NewToken(TkString, position, 0, str)
	cur.Next = tok
	return tok
}

func NewTokenFloat(cur *Token, position *Position, num float64, str string) *Token {
	tok := NewToken(TkFloat, position, num, str)
	cur.Next = tok
	return tok
}

func NewTokenInt(cur *Token, position *Position, num float64, str string) *Token {
	tok := NewToken(TkInt, position, num, str)
	cur.Next = tok
	return tok
}
