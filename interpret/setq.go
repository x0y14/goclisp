package interpret

import (
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func setq(storage *Storage, node *parse.Node) (*data.Data, error) {
	if storage == nil {
		storage = globalVariables
	}
	if len(node.Arguments)%2 != 0 {
		return nil, NewRuntimeError(AssignErr, "the number of key and value does not matched")
	}
	i := 0
	for i < len(node.Arguments) {
		key := node.Arguments[i]
		i++
		value := node.Arguments[i]
		i++
		err := storeData(storage, key.Value.Atom.Str, value.Value)
		if err != nil {
			return nil, err
		}
	}
	// return last one
	return node.Arguments[i-1].Value, nil
}
