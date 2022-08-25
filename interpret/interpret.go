package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/atom"
	"github.com/x0y14/goclisp/parse"
)

func add(node *parse.Node) (*atom.Atom, error) {
	floatMode := false
	var result float64 = 0

	for i, arg := range node.Arguments {
		var diff float64
		var diffKind atom.Kind
		switch arg.Kind {
		case parse.Float, parse.Int:
			diff = arg.Value.Num
			diffKind = arg.Value.Kind
		case parse.String, parse.True, parse.Nil:
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
		case parse.Ident:
			// todo
		default:
			a, err := exec(arg)
			if err != nil {
				return nil, err
			}
			diff = a.Num
			diffKind = a.Kind
		}
		if diffKind == atom.Float {
			floatMode = true
		}

		if i == 0 {
			result = diff
		} else {
			result += diff
		}
	}

	if floatMode {
		return atom.NewAtomF(result), nil
	}
	return atom.NewAtomI(result), nil
}

func exec(node *parse.Node) (*atom.Atom, error) {
	switch node.Kind {
	// atom
	case parse.String, parse.Float, parse.Int, parse.True, parse.Nil:
		return node.Value, nil
	// arithmetic op
	case parse.Add:
		return add(node)
	}

	return nil, NewRuntimeError(
		UnimplementedErr,
		fmt.Sprintf("unimplemented: NodeKind(%d)", node.Kind))
}

func Interpret(nodes []*parse.Node) error {
	for _, node := range nodes {
		v, err := exec(node)
		if err != nil {
			return nil
		}
		fmt.Println(v.String())
	}
	return nil
}
