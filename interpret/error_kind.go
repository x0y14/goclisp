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
	FunctionArgumentErr
)

var errorKind = [...]string{
	UnknownErr:          "UnknownErr",
	TypeMissMatchErr:    "TypeMissMatchErr",
	UndefinedErr:        "UndefinedErr",
	UnimplementedErr:    "Unimplemented",
	DivideByZeroErr:     "DivideByZero",
	AssignErr:           "AssignErr",
	FunctionArgumentErr: "FunctionArgumentErr",
}

func (r RuntimeErrorKind) String() string {
	return errorKind[r]
}
