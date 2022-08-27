package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
)

func exec(node *parse.Node) (*data.Data, error) {
	switch node.Kind {
	// atom
	case parse.String, parse.Float, parse.Int, parse.True, parse.Nil:
		return node.Value, nil
	// arithmetic op
	case parse.Add, parse.Sub, parse.Mul, parse.Div:
		return addSubMulDiv(node)
	// logical op
	case parse.Eq, parse.Ne:
		return eqNe(node)
	case parse.Lt, parse.Le, parse.Gt, parse.Ge:
		return ltLeGtGe(node)
	case parse.Ident:
		return globalLoad(node.Value.Atom.Str)
	case parse.Call:
		switch node.Value.Atom.Str {
		case "format":
			return format(node)
		case "setq":
			return setq(node)
		}
	}

	return nil, NewRuntimeError(
		UnimplementedErr,
		fmt.Sprintf("unimplemented: NodeKind(%d)", node.Kind))
}

func Interpret(nodes []*parse.Node) error {
	for _, node := range nodes {
		v, err := exec(node)
		if err != nil {
			return err
		}
		fmt.Println(v.String())
	}
	return nil
}
