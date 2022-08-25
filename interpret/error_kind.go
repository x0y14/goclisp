package interpret

type RuntimeErrorKind int

const (
	_ RuntimeErrorKind = iota
	UnknownErr
	TypeMissMatchErr
	UndefinedErr
	UnimplementedErr
)

var errorKind = [...]string{
	UnknownErr:       "UnknownErr",
	TypeMissMatchErr: "TypeMissMatchErr",
	UndefinedErr:     "UndefinedErr",
	UnimplementedErr: "Unimplemented",
}

func (r RuntimeErrorKind) String() string {
	return errorKind[r]
}
