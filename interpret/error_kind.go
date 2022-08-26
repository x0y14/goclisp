package interpret

type RuntimeErrorKind int

const (
	_ RuntimeErrorKind = iota
	UnknownErr
	TypeMissMatchErr
	UndefinedErr
	UnimplementedErr
	DivideByZeroErr
	AssignErr
)

var errorKind = [...]string{
	UnknownErr:       "UnknownErr",
	TypeMissMatchErr: "TypeMissMatchErr",
	UndefinedErr:     "UndefinedErr",
	UnimplementedErr: "Unimplemented",
	DivideByZeroErr:  "DivideByZero",
	AssignErr:        "AssignErr",
}

func (r RuntimeErrorKind) String() string {
	return errorKind[r]
}
