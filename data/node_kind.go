package data

type NodeKind int

const (
	_ NodeKind = iota

	atomBegin
	NdString // "abc"
	NdFloat  // 1.0
	NdInt    // 1
	NdTrue   // T
	NdNil    // NIL
	NdIdent  //
	atomEnd

	arithmeticOperatorBegin
	NdAdd // +
	NdSub // -
	NdMul // *
	NdDiv // /
	arithmeticOperatorEnd

	logicalOperatorBegin
	NdEq // =
	NdNe // /=
	NdLt // <
	NdLe // <=
	NdGt // >
	NdGe // >=
	logicalOperatorEnd

	NdCall
)
