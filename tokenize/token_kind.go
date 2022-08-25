package tokenize

type TokenKind int

const (
	_ TokenKind = iota
	Eof
	Reserved
	Ident
	String
	Number
)
