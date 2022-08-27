package interpret

import (
	"github.com/x0y14/goclisp/data"
)

func ltLeGtGe(node *data.Node) (*data.Data, error) {
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

		if node.Kind == data.NdLt && !(base < v.Atom.Num) {
			return data.NewDataNil(), nil
		}

		if node.Kind == data.NdLe && !(base <= v.Atom.Num) {
			return data.NewDataNil(), nil
		}

		if node.Kind == data.NdGt && !(base > v.Atom.Num) {
			return data.NewDataNil(), nil
		}

		if node.Kind == data.NdGe && !(base >= v.Atom.Num) {
			return data.NewDataNil(), nil
		}
	}

	return data.NewDataTrue(), nil
}
