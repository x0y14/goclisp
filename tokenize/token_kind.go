package tokenize

type TokenKind int

const (
	Illegal TokenKind = iota
	Eof
	Reserved
	Ident
	String
	Number
)
