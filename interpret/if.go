package interpret

import "github.com/x0y14/goclisp/data"

func if_(scope *data.Storage, node *data.Node) (*data.Data, error) {
	if len(node.Arguments) != 3 {
		return nil, NewRuntimeError(AssignErr, "need ( cond, true-block, false-block )")
	}

	// 条件式を検証
	tf, err := eval(scope, node.Arguments[0])
	if err != nil {
		return nil, err
	}

	// trueだった場合
	if tf.IsAtom() && tf.Atom.Kind == data.True {
		return eval(scope, node.Arguments[1])
	}
	// falseだった場合
	return eval(scope, node.Arguments[2])
}
