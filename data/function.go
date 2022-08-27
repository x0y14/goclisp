package data

type Function struct {
	Params       []*Atom
	Description  string
	Body         []*Node
	LocalStorage *Storage
}

func NewFunction(params []*Atom, desc string, body []*Node) *Function {
	return &Function{
		Params:       params,
		Description:  desc,
		Body:         body,
		LocalStorage: NewStorage(),
	}
}
