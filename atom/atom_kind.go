package atom

type Kind int

const (
	_ Kind = iota
	String
	Float
	Int
	True
	Nil
	Ident
)
