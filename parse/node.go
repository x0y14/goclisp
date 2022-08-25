package parse

type Node struct {
	Kind      NodeKind
	Arguments []*Node

	Num float64 // int, floatの場合に値が入る
	Str string  // ident, stringの場合に値が入る
}

func NewNodeString(str string) *Node {
	return &Node{
		Kind: String,
		Str:  str,
	}
}

func NewNodeFloat(num float64) *Node {
	return &Node{
		Kind: Float,
		Num:  num,
	}
}

func NewNodeInt(num float64) *Node {
	return &Node{
		Kind: Int,
		Num:  num,
	}
}

func NewNodeTrue() *Node {
	return &Node{
		Kind: True,
	}
}

func NewNodeNil() *Node {
	return &Node{
		Kind: Nil,
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		Kind: Ident,
		Str:  ident,
	}
}

func NewNodeWithArgs(kind NodeKind, args []*Node) *Node {
	return &Node{
		Kind:      kind,
		Arguments: args,
	}
}
