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

		val, err := eval(arg)
		if err != nil {
			return nil, err
		}
		switch val.Atom.Kind {
		case data.Float:
			diff = val.Atom.Num
			floatMode = true
		case data.Int:
			diff = val.Atom.Num
		default:
			return nil, NewRuntimeError(TypeMissMatchErr, "the value type is not float or int")
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
