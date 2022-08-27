package interpret

import (
	"github.com/x0y14/goclisp/data"
)

func defun(scope *data.Storage, node *data.Node) (*data.Data, error) {
	// 第一引数がidentと識別された場合、パースされたのち、Valueに入る
	// その他は全てArgumentsに入っているので、解析に必要なatomだけ取り出し、全てまとめてparamsとする

	// node.value: "defun"
	// node.arguments
	//   0: func-name
	//   1: param-list
	//  2?: optional comment
	//2?3?: body

	// setq同様、新しく定義する関数なのでevalは使わないで直接値を取り出す。

	nameNode := node.Arguments[0]
	paramNode := node.Arguments[1]

	name := nameNode.Value.Atom.Str

	params := []*data.Atom{paramNode.Value.Atom}
	for _, arg := range paramNode.Arguments {
		params = append(params, arg.Value.Atom)
	}

	var desc string
	var body []*data.Node

	if node.Arguments[2].Kind == data.NdString {
		// has optional description
		desc = node.Arguments[2].Value.Atom.Str
		body = node.Arguments[3:] // 3以降
	} else {
		desc = ""
		body = node.Arguments[2:] // 2以降
	}

	f := data.NewFunction(params, desc, body)
	err := data.StoreData(scope, name, data.NewDataFunc(f))
	if err != nil {
		return nil, err
	}

	// identなのでstringしても内容がないよう
	return nameNode.Value, nil
}
