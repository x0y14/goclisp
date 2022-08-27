package interpret

import (
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func eqNe(node *parse.Node) (*data.Data, error) {
	var base float64
	for i, arg := range node.Arguments {
		// type check
		if arg.Kind != parse.Float && arg.Kind != parse.Int {
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
		}

		if i == 0 {
			base = arg.Value.Atom.Num
			continue
		}

		if node.Kind == parse.Eq && base != arg.Value.Atom.Num {
			return data.NewDataNil(), nil
		}

		if node.Kind == parse.Ne && base == arg.Value.Atom.Num {
			return data.NewDataNil(), nil
		}
	}

	return data.NewDataTrue(), nil
}
