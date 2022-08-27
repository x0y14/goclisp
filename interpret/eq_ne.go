package interpret

import (
	"github.com/x0y14/goclisp/data"
)

func eqNe(node *data.Node) (*data.Data, error) {
	var base float64
	for i, arg := range node.Arguments {
		v, err := eval(nil, arg)
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

		if node.Kind == data.NdEq && base != v.Atom.Num {
			return data.NewDataNil(), nil
		}

		if node.Kind == data.NdNe && base == v.Atom.Num {
			return data.NewDataNil(), nil
		}
	}

	return data.NewDataTrue(), nil
}
