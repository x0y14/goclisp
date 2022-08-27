package interpret

import (
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func eqNe(node *parse.Node) (*data.Data, error) {
	var base float64
	for i, arg := range node.Arguments {
		v, err := eval(arg)
		if err != nil {
			return nil, err
		}

		if !v.IsAtom() || (v.Atom.Kind != data.Float && v.Atom.Kind != data.Int) {
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
		}

		if i == 0 {
			base = v.Atom.Num
			continue
		}

		if node.Kind == parse.Eq && base != v.Atom.Num {
			return data.NewDataNil(), nil
		}

		if node.Kind == parse.Ne && base == v.Atom.Num {
			return data.NewDataNil(), nil
		}
	}

	return data.NewDataTrue(), nil
}
