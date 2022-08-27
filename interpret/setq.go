package interpret

import (
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func setq(storage *data.Storage, node *parse.Node) (*data.Data, error) {
	if storage == nil {
		storage = globalVariables
	}
	if len(node.Arguments)%2 != 0 {
		return nil, NewRuntimeError(AssignErr, "the number of key and value does not matched")
	}
	i := 0
	var result *data.Data

	for i < len(node.Arguments) {
		// 評価しない
		key := node.Arguments[i]
		i++

		// 評価する
		value := node.Arguments[i]
		v, err := eval(value)
		if err != nil {
			return nil, err
		}
		i++

		err = data.StoreData(storage, key.Value.Atom.Str, v)
		if err != nil {
			return nil, err
		}
		result = v
	}
	// return last one
	return result, nil
}
