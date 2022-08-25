package interpret

type RuntimeErrorKind int

const (
	_ RuntimeErrorKind = iota
	UnknownErr
	TypeMissMatchErr
	UndefinedErr
	UnimplementedErr
	DivideByZeroErr
)

var errorKind = [...]string{
	UnknownErr:       "UnknownErr",
	TypeMissMatchErr: "TypeMissMatchErr",
	UndefinedErr:     "UndefinedErr",
	UnimplementedErr: "Unimplemented",
	DivideByZeroErr:  "DivideByZero",
}

func (r RuntimeErrorKind) String() string {
	return errorKind[r]
}
