package atom

import "fmt"

type Atom struct {
	Kind Kind

	Num float64
	//NumI int
	Str string
}

func (a *Atom) String() string {
	switch a.Kind {
	case String:
		return a.Str
	case Float:
		return fmt.Sprintf("%f", a.Num)
	case Int:
		return fmt.Sprintf("%d", int(a.Num))
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
		Num:  num,
	}
}

func NewAtomI(num float64) *Atom {
	return &Atom{
		Kind: Int,
		Num:  num,
	}
}

func NewAtomTrue() *Atom {
	return &Atom{Kind: True}
}

func NewAtomNil() *Atom {
	return &Atom{Kind: Nil}
}

func NewAtomIdent(str string) *Atom {
	return &Atom{Kind: Ident, Str: str}
}
