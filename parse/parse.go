package parse

import (
	"github.com/x0y14/goclisp/data"
)

var token *data.Token

func consumeReserved(str string) bool {
	if token.Kind != data.TkReserved || token.Str != str {
		return false
	}
	token = token.Next
	return true
}

func consumeIdent() *data.Token {
	if token.Kind != data.TkIdent {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeString() *data.Token {
	if token.Kind != data.TkString {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeFloat() *data.Token {
	if token.Kind != data.TkFloat {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeInt() *data.Token {
	if token.Kind != data.TkInt {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeNil() *data.Token {
	if token.Kind != data.TkIdent || (token.Str != "NIL" && token.Str != "nil") {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeTrue() *data.Token {
	if token.Kind != data.TkIdent || (token.Str != "T" && token.Str != "t") {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func expectOpIdent() (*data.Token, error) {
	if token.Kind != data.TkReserved && token.Kind != data.TkIdent {
		return nil, NewSyntaxError("unexpected token", token)
	}
	tok := token
	token = token.Next
	return tok, nil
}

func atEof() bool {
	return token.Kind == data.TkEof
}

func program() ([]*data.Node, error) {
	var stmts []*data.Node
	for !atEof() {
		s, err := stmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, s)
	}
	return stmts, nil
}

func stmt() (*data.Node, error) {
	if consumeReserved("(") {
		oi, err := expectOpIdent()
		if err != nil {
			return nil, err
		}

		var args []*data.Node
		for !consumeReserved(")") {
			arg, err := stmt()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}

		if oi.Kind == data.TkIdent {
			return data.NewNodeCall(oi.Str, args), nil
		}
		switch oi.Str {
		case "+":
			return data.NewNodeWithArgs(data.NdAdd, args), nil
		case "-":
			return data.NewNodeWithArgs(data.NdSub, args), nil
		case "*":
			return data.NewNodeWithArgs(data.NdMul, args), nil
		case "/":
			return data.NewNodeWithArgs(data.NdDiv, args), nil
		case "=":
			return data.NewNodeWithArgs(data.NdEq, args), nil
		case "/=":
			return data.NewNodeWithArgs(data.NdNe, args), nil
		case "<":
			return data.NewNodeWithArgs(data.NdLt, args), nil
		case "<=":
			return data.NewNodeWithArgs(data.NdLe, args), nil
		case ">":
			return data.NewNodeWithArgs(data.NdGt, args), nil
		case ">=":
			return data.NewNodeWithArgs(data.NdGe, args), nil
		}

	}
	return unary()
}

func unary() (*data.Node, error) {
	if consumeReserved("+") {
		return primary()
	} else if consumeReserved("-") {
		p, err := primary()
		if err != nil {
			return nil, err
		}
		// floatとintが同じ式に入っていたら、floatが優先される(intがfloatにキャストされる)
		// 1 + 2.0 = 3.0
		// なので、int(0)を使えば、pがintの場合でも、floatの場合でもint(0)がキャストされるだけでpの型が変更されることはない
		return data.NewNodeWithArgs(data.NdSub, []*data.Node{data.NewNodeInt(0), p}), nil
	}
	return primary()
}

func primary() (*data.Node, error) {
	var p *data.Token
	if p = consumeString(); p != nil {
		return data.NewNodeString(p.Str), nil
	} else if p = consumeFloat(); p != nil {
		return data.NewNodeFloat(p.Num), nil
	} else if p = consumeInt(); p != nil {
		return data.NewNodeInt(p.Num), nil
	} else if p = consumeNil(); p != nil {
		return data.NewNodeNil(), nil
	} else if p = consumeTrue(); p != nil {
		return data.NewNodeTrue(), nil
	} else if p = consumeIdent(); p != nil {
		return data.NewNodeIdent(p.Str), nil
	}

	return nil, NewSyntaxError("unexpected token", token)
}

func Parse(tok *data.Token) ([]*data.Node, error) {
	token = tok
	return program()
}
