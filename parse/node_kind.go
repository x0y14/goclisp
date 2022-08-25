package parse

type NodeKind int

const (
	_ NodeKind = iota

	atomBegin
	String // "abc"
	Float  // 1.0
	Int    // 1
	True   // T
	Nil    // NIL
	Ident  //
	atomEnd

	arithmeticOperatorBegin
	Add // +
	Sub // -
	Mul // *
	Div // /
	arithmeticOperatorEnd

	logicalOperatorBegin
	Eq // =
	Ne // /=
	Lt // <
	Le // <=
	Gt // >
	Ge // >=
	logicalOperatorEnd

	Call
)
