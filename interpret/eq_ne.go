package interpret

import (
	"github.com/x0y14/goclisp/atom"
	"github.com/x0y14/goclisp/parse"
)

func eqNe(kind parse.NodeKind, node *parse.Node) (*atom.Atom, error) {
	var base float64
	for i, arg := range node.Arguments {
		// type check
		if arg.Kind != parse.Float && arg.Kind != parse.Int {
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
		}

		if i == 0 {
			base = arg.Value.Num
			continue
		}

		if kind == parse.Eq && base != arg.Value.Num {
			return atom.NewAtomNil(), nil
		}

		if kind == parse.Ne && base == arg.Value.Num {
			return atom.NewAtomNil(), nil
		}
	}

	return atom.NewAtomTrue(), nil
}
