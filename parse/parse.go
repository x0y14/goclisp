package parse

import "github.com/x0y14/goclisp/tokenize"

var token *tokenize.Token

func consumeReserved(str string) bool {
	if token.Kind != tokenize.Reserved || token.Str != str {
		return false
	}
	token = token.Next
	return true
}

func consumeIdent() *tokenize.Token {
	if token.Kind != tokenize.Ident {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeString() *tokenize.Token {
	if token.Kind != tokenize.String {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeFloat() *tokenize.Token {
	if token.Kind != tokenize.Float {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeInt() *tokenize.Token {
	if token.Kind != tokenize.Int {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeNil() *tokenize.Token {
	if token.Kind != tokenize.Ident || (token.Str != "NIL" && token.Str != "nil") {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func consumeTrue() *tokenize.Token {
	if token.Kind != tokenize.Ident || (token.Str != "T" && token.Str != "t") {
		return nil
	}
	tok := token
	token = token.Next
	return tok
}

func expectOpIdent() (*tokenize.Token, error) {
	if token.Kind != tokenize.Reserved && token.Kind != tokenize.Ident {
		return nil, NewSyntaxError("unexpected token", token)
	}
	tok := token
	token = token.Next
	return tok, nil
}

func atEof() bool {
	return token.Kind == tokenize.Eof
}

func program() ([]*Node, error) {
	var stmts []*Node
	for !atEof() {
		s, err := stmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, s)
	}
	return stmts, nil
}

func stmt() (*Node, error) {
	if consumeReserved("(") {
		oi, err := expectOpIdent()
		if err != nil {
			return nil, err
		}

		var args []*Node
		for !consumeReserved(")") {
			arg, err := stmt()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}

		if oi.Kind == tokenize.Ident {
			return NewNodeCall(oi.Str, args), nil
		}
		switch oi.Str {
		case "+":
			return NewNodeWithArgs(Add, args), nil
		case "-":
			return NewNodeWithArgs(Sub, args), nil
		case "*":
			return NewNodeWithArgs(Mul, args), nil
		case "/":
			return NewNodeWithArgs(Div, args), nil
		case "=":
			return NewNodeWithArgs(Eq, args), nil
		case "/=":
			return NewNodeWithArgs(Ne, args), nil
		case "<":
			return NewNodeWithArgs(Lt, args), nil
		case "<=":
			return NewNodeWithArgs(Le, args), nil
		case ">":
			return NewNodeWithArgs(Gt, args), nil
		case ">=":
			return NewNodeWithArgs(Ge, args), nil
		}

	}
	return unary()
}

func unary() (*Node, error) {
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
		return NewNodeWithArgs(Sub, []*Node{NewNodeInt(0), p}), nil
	}
	return primary()
}

func primary() (*Node, error) {
	var p *tokenize.Token
	if p = consumeString(); p != nil {
		return NewNodeString(p.Str), nil
	} else if p = consumeFloat(); p != nil {
		return NewNodeFloat(p.Num), nil
	} else if p = consumeInt(); p != nil {
		return NewNodeInt(p.Num), nil
	} else if p = consumeNil(); p != nil {
		return NewNodeNil(), nil
	} else if p = consumeTrue(); p != nil {
		return NewNodeTrue(), nil
	} else if p = consumeIdent(); p != nil {
		return NewNodeIdent(p.Str), nil
	}

	return nil, NewSyntaxError("unexpected token", token)
}

func Parse(tok *tokenize.Token) ([]*Node, error) {
	token = tok
	return program()
}
