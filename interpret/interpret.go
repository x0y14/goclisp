package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/parse"
)

func exec(node *parse.Node) (*Atom, error) {
	switch node.Kind {
	case parse.String:
		return NewAtomString(node.Str), nil
	case parse.Float:
		return NewAtomF(node.Num), nil
	case parse.Int:
		return NewAtomI(int(node.Num)), nil
	case parse.True:
		return NewAtomTrue(), nil
	case parse.Nil:
		return NewAtomNil(), nil
	}

	return nil, NewRuntimeError(
		UnimplementedErr,
		fmt.Sprintf("unimplemented: NodeKind(%d)", node.Kind))
}

func Interpret(nodes []*parse.Node) error {
	for _, node := range nodes {
		atom, err := exec(node)
		if err != nil {
			return nil
		}
		fmt.Println(atom.String())
	}
	return nil
}
