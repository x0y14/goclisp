package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/atom"
	"github.com/x0y14/goclisp/parse"
)

func exec(node *parse.Node) (*atom.Atom, error) {
	switch node.Kind {
	// atom
	case parse.String, parse.Float, parse.Int, parse.True, parse.Nil:
		return node.Value, nil
	// arithmetic op
	case parse.Add, parse.Sub, parse.Mul, parse.Div:
		return addSubMulDiv(node.Kind, node)
	}

	return nil, NewRuntimeError(
		UnimplementedErr,
		fmt.Sprintf("unimplemented: NodeKind(%d)", node.Kind))
}

func Interpret(nodes []*parse.Node) error {
	for _, node := range nodes {
		v, err := exec(node)
		if err != nil {
			return nil
		}
		fmt.Println(v.String())
	}
	return nil
}
