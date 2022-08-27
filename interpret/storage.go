package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
)

type Storage map[string]*data.Data

func NewStorage() *Storage {
	return &Storage{}
}

func storeData(storage *Storage, key string, value *data.Data) error {
	if value.Kind == data.Atomic && value.Atom.Kind == data.Ident {
		return NewRuntimeError(AssignErr, "can't assign ident to ident")
	}
	(*storage)[key] = value
	return nil
}

func loadData(storage *Storage, key string) (*data.Data, error) {
	v, ok := (*storage)[key]
	if !ok {
		return nil, NewRuntimeError(UndefinedErr, fmt.Sprintf("%s is not defined", key))
	}
	return v, nil
}

func deleteData(storage *Storage, key string) (*data.Data, error) {
	v, err := loadData(storage, key)
	if err != nil {
		return nil, err
	}
	delete(*storage, key)
	return v, nil
}
