package interpret

import (
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func addSubMulDiv(node *parse.Node) (*data.Data, error) {
	floatMode := false
	var result float64 = 0

	for i, arg := range node.Arguments {
		var diff float64
		var diffKind data.AtomKind
		switch arg.Kind {
		case parse.Float, parse.Int:
			diff = arg.Value.Atom.Num
			diffKind = arg.Value.Atom.Kind
		case parse.String, parse.True, parse.Nil:
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
		case parse.Ident:
			// todo
		default:
			a, err := exec(arg)
			if err != nil {
				return nil, err
			}
			diff = a.Atom.Num
			diffKind = a.Atom.Kind
		}
		if diffKind == data.Float {
			floatMode = true
		}

		if i == 0 {
			result = diff
		} else {
			switch node.Kind {
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
		return data.NewDataFloat(result), nil
	}
	return data.NewDataInt(result), nil
}
