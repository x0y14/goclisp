package parse

import (
	"github.com/x0y14/goclisp/data"
)

type Node struct {
	Kind      NodeKind
	Arguments []*Node
	Value     *data.Data
}

func NewNodeString(str string) *Node {
	return &Node{
		Kind:  String,
		Value: data.NewDataString(str),
	}
}

func NewNodeFloat(num float64) *Node {
	return &Node{
		Kind:  Float,
		Value: data.NewDataFloat(num),
	}
}

func NewNodeInt(num float64) *Node {
	return &Node{
		Kind:  Int,
		Value: data.NewDataInt(num),
	}
}

func NewNodeTrue() *Node {
	return &Node{
		Kind:  True,
		Value: data.NewDataTrue(),
	}
}

func NewNodeNil() *Node {
	return &Node{
		Kind:  Nil,
		Value: data.NewDataNil(),
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		Kind:  Ident,
		Value: data.NewDataIdent(ident),
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
		Kind:      Call,
		Arguments: args,
		Value:     data.NewDataIdent(ident),
	}
}
