package data

type AtomKind int

const (
	_ AtomKind = iota
	String
	Float
	Int
	True
	Nil
	Ident
)
