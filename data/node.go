package data

type Node struct {
	Kind      NodeKind
	Arguments []*Node
	Value     *Data
}

func NewNodeString(str string) *Node {
	return &Node{
		Kind:  NdString,
		Value: NewDataString(str),
	}
}

func NewNodeFloat(num float64) *Node {
	return &Node{
		Kind:  NdFloat,
		Value: NewDataFloat(num),
	}
}

func NewNodeInt(num float64) *Node {
	return &Node{
		Kind:  NdInt,
		Value: NewDataInt(num),
	}
}

func NewNodeTrue() *Node {
	return &Node{
		Kind:  NdTrue,
		Value: NewDataTrue(),
	}
}

func NewNodeNil() *Node {
	return &Node{
		Kind:  NdNil,
		Value: NewDataNil(),
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		Kind:  NdIdent,
		Value: NewDataIdent(ident),
	}
}

func NewNodeWithArgs(kind NodeKind, args []*Node) *Node {
	return &Node{
		Kind:      kind,
		Arguments: args,
	}
}

func NewNodeCall(ident string, args []*Node) *Node {
	return &Node{
		Kind:      NdCall,
		Arguments: args,
		Value:     NewDataIdent(ident),
	}
}
