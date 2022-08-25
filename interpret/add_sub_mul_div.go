package interpret

import (
	"github.com/x0y14/goclisp/atom"
	"github.com/x0y14/goclisp/parse"
)

func addSubMulDiv(kind parse.NodeKind, node *parse.Node) (*atom.Atom, error) {
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
			switch kind {
			case parse.Add:
				result += diff
			case parse.Sub:
				result -= diff
			case parse.Mul:
				result *= diff
			case parse.Div:
				if diff == 0 {
					return nil, NewRuntimeError(DivideByZeroErr, "division by zero")
				}
				result /= diff
			}
		}
	}

	if floatMode {
		return atom.NewAtomF(result), nil
	}
	return atom.NewAtomI(result), nil
}
