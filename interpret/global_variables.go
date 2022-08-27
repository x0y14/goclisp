package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
)

var globalVariables map[string]*data.Data

func init() {
	globalVariables = map[string]*data.Data{}
}

func globalStore(key string, value *data.Data) error {
	if value.Kind == data.Atomic && value.Atom.Kind == data.Ident {
		return NewRuntimeError(AssignErr, "can't assign ident to ident")
	}
	globalVariables[key] = value
	return nil
}

func globalLoad(key string) (*data.Data, error) {
	v, ok := globalVariables[key]
	if !ok {
		return nil, NewRuntimeError(UndefinedErr, fmt.Sprintf("%s is not defined", key))
	}
	return v, nil
}
