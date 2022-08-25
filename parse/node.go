package parse

import "github.com/x0y14/goclisp/atom"

type Node struct {
	Kind      NodeKind
	Arguments []*Node
	Value     *atom.Atom
}

func NewNodeString(str string) *Node {
	return &Node{
		Kind:  String,
		Value: atom.NewAtomString(str),
	}
}

func NewNodeFloat(num float64) *Node {
	return &Node{
		Kind:  Float,
		Value: atom.NewAtomF(num),
	}
}

func NewNodeInt(num float64) *Node {
	return &Node{
		Kind:  Int,
		Value: atom.NewAtomI(num),
	}
}

func NewNodeTrue() *Node {
	return &Node{
		Kind:  True,
		Value: atom.NewAtomTrue(),
	}
}

func NewNodeNil() *Node {
	return &Node{
		Kind:  Nil,
		Value: atom.NewAtomNil(),
	}
}

func NewNodeIdent(ident string) *Node {
	return &Node{
		Kind:  Ident,
		Value: atom.NewAtomIdent(ident),
	}
}

func NewNodeWithArgs(kind NodeKind, args []*Node) *Node {
	return &Node{
		Kind:      kind,
		Arguments: args,
	}
}
