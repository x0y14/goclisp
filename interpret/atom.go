package interpret

import "fmt"

type Atom struct {
	Kind AtomKind

	NumF float64
	NumI int
	Str  string
}

func (a *Atom) String() string {
	switch a.Kind {
	case String:
		return a.Str
	case Float:
		return fmt.Sprintf("%f", a.NumF)
	case Int:
		return fmt.Sprintf("%d", a.NumI)
	case True:
		return "T"
	case Nil:
		return "NIL"
	}
	return ""
}

func NewAtomString(str string) *Atom {
	return &Atom{
		Kind: String,
		Str:  str,
	}
}

func NewAtomF(num float64) *Atom {
	return &Atom{
		Kind: Float,
		NumF: num,
	}
}

func NewAtomI(num int) *Atom {
	return &Atom{
		Kind: Int,
		NumI: num,
	}
}

func NewAtomTrue() *Atom {
	return &Atom{Kind: True}
}

func NewAtomNil() *Atom {
	return &Atom{Kind: Nil}
}
