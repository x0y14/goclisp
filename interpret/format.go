package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
	"strings"
)

func format(scope *data.Storage, node *data.Node) (*data.Data, error) {
	printMode := true
	// arg[0] = true
	// arg[1] = format
	// arg[2:]= values
	if len(node.Arguments) <= 1 {
		return nil, NewRuntimeError(FunctionArgumentErr, "invalid number of argument")
	}
	if node.Arguments[0].Kind != data.NdTrue && node.Arguments[0].Kind != data.NdNil {
		return nil, NewRuntimeError(UnimplementedErr, "format support T, NIL")
	}
	if node.Arguments[0].Kind == data.NdNil {
		printMode = false
	}

	// cg "%s"
	if node.Arguments[1].Value.Atom.Kind != data.String {
		return nil, NewRuntimeError(FunctionArgumentErr, "need format-str")
	}
	f := node.Arguments[1].Value.Atom.Str
	f = strings.ReplaceAll(f, "~A", "%s")

	var args []any
	i := 2
	// t "%s" ?
	if len(node.Arguments) >= 3 {
		for i < len(node.Arguments) {
			a, err := eval(scope, node.Arguments[i])
			if err != nil {
				return nil, err
			}
			args = append(args, a.String())
			i++
		}
	}

	if printMode {
		fmt.Printf(f+"\n", args...)
		return data.NewDataNil(), nil
	}
	return data.NewDataString(fmt.Sprintf(f, args...)), nil
}
