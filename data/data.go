package data

type Data struct {
	Kind
	Atom     *Atom
	Function *Function
}

func (d *Data) String() string {
	switch d.Kind {
	case Atomic:
		return d.Atom.String()
	}
	panic("data is not printable")
}

func (d *Data) IsAtom() bool {
	return d.Kind == Atomic
}

func newAtomic(atom *Atom) *Data {
	return &Data{
		Kind:     Atomic,
		Atom:     atom,
		Function: nil,
	}
}

func NewDataString(str string) *Data {
	return newAtomic(NewAtomString(str))
}

func NewDataFloat(num float64) *Data {
	return newAtomic(NewAtomF(num))
}

func NewDataInt(num float64) *Data {
	return newAtomic(NewAtomI(num))
}

func NewDataTrue() *Data {
	return newAtomic(NewAtomTrue())
}

func NewDataNil() *Data {
	return newAtomic(NewAtomNil())
}

func NewDataIdent(str string) *Data {
	return newAtomic(NewAtomIdent(str))
}

func NewDataFunc(f *Function) *Data {
	return &Data{
		Kind:     Fn,
		Atom:     nil,
		Function: f,
	}
}
