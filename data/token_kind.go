package data

type TokenKind int

const (
	_ TokenKind = iota
	TkEof
	TkReserved
	TkIdent
	TkString
	TkInt
	TkFloat
)
