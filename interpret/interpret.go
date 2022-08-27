package interpret

import (
	"fmt"
	"github.com/x0y14/goclisp/data"
)

func call(node *data.Node, f *data.Function) (*data.Data, error) {
	var result *data.Data

	// 関数呼び出しに際して、引数の受け渡しをする
	// LocalStorageにNode.Argumentsをぶち込む

	// 期待していた引数の数と渡された引数の数が一致しない
	if len(f.Params) != len(node.Arguments) {
		return nil, NewRuntimeError(FunctionArgumentErr, "invalid number of argument")
	}

	// f.LocalStorage[f.Param[i]] = node.Arguments[i]
	for i, param := range f.Params {
		// paramはローカル変数とする名前
		err := data.StoreData(f.LocalStorage, param.Str, node.Arguments[i].Value)
		if err != nil {
			return nil, err
		}
	}

	// 各ノード実行していく
	for _, b := range f.Body {
		// スコープは関数ローカルストレージにしてあげる
		d, err := eval(f.LocalStorage, b)
		if err != nil {
			return nil, err
		}
		result = d
	}

	// 最終ノードの結果を返してあげる
	return result, nil
}

func eval(scope *data.Storage, node *data.Node) (*data.Data, error) {
	//if scope == nil {
	//	scope = GlobalVariables
	//}

	switch node.Kind {
	// atom
	case data.NdString, data.NdFloat, data.NdInt, data.NdTrue, data.NdNil:
		return node.Value, nil
	// arithmetic op
	case data.NdAdd, data.NdSub, data.NdMul, data.NdDiv:
		return addSubMulDiv(scope, node)
	// logical op
	case data.NdEq, data.NdNe:
		return eqNe(node)
	case data.NdLt, data.NdLe, data.NdGt, data.NdGe:
		return ltLeGtGe(node)
	case data.NdIdent:
		return data.LoadData(scope, node.Value.Atom.Str)
	case data.NdCall:
		switch node.Value.Atom.Str {
		case "format":
			return format(scope, node)
		case "setq":
			return setq(scope, node)
		case "defun":
			name := node.Arguments[0].Value.Atom.Str
			_, err := data.LoadData(scope, name)
			if err == nil {
				return nil, NewRuntimeError(AlreadyDefinedErr, fmt.Sprintf("%s is defined", name))
			}
			return defun(scope, node)
		default:
			name := node.Value.Atom.Str
			f, err := data.LoadData(scope, name)
			if err != nil {
				return nil, err
			}
			return call(node, f.Function)
		}
	}

	return nil, NewRuntimeError(
		UnimplementedErr,
		fmt.Sprintf("unimplemented: NodeKind(%d)", node.Kind))
}

func Interpret(nodes []*data.Node) error {
	for _, node := range nodes {
		v, err := eval(data.GlobalStorage, node)
		if err != nil {
			return err
		}
		fmt.Println(v.String())
	}
	return nil
}
