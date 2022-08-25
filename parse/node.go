package parse

type Node struct {
	kind      NodeKind
	arguments []*Node

	num float64 // int, floatの場合に値が入る
	str string  // ident, stringの場合に値が入る
}

func NewNodeString(str string) *Node {
	return &Node{
		kind: String,
		str:  str,
	}
}

func NewNodeFloat(num float64) *Node {
	return &Node{
		kind: Float,
		num:  num,
	}
}

func NewNodeInt(num float64) *Node {
	return &Node{
		kind: Int,
		num:  num,
	}
}

func NewNodeTrue() *Node {
	return &Node{
		kind: True,
	}
}

func NewNodeNil() *Node {
	return &Node{
		kind: Nil,
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		kind: Ident,
		str:  ident,
	}
}

func NewNodeWithArgs(kind NodeKind, args []*Node) *Node {
	return &Node{
		kind:      kind,
		arguments: args,
	}
}
