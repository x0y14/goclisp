package data

import (
	"fmt"
)

var GlobalStorage *Storage

func init() {
	GlobalStorage = NewStorage()
}

type Storage map[string]*Data

func NewStorage() *Storage {
	return &Storage{}
}

func StoreData(storage *Storage, key string, value *Data) error {
	if value.Kind == Atomic && value.Atom.Kind == Ident {
		return fmt.Errorf("can't assign ident to ident")
	}
	(*storage)[key] = value
	return nil
}

func LoadData(storage *Storage, key string) (*Data, error) {
	v, ok := (*storage)[key]
	if !ok {
		gV, gOk := (*GlobalStorage)[key]
		if gOk {
			return gV, nil
		}
		return nil, fmt.Errorf("%s is not defined", key)
	}
	return v, nil
}

func DeleteData(storage *Storage, key string) (*Data, error) {
	v, err := LoadData(storage, key)
	if err != nil {
		return nil, err
	}
	delete(*storage, key)
	return v, nil
}
