package interpret

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/atom"
	"github.com/x0y14/goclisp/parse"
	"github.com/x0y14/goclisp/tokenize"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*atom.Atom
	}{
		// atom
		{
			"string",
			"\"hello\"",
			[]*atom.Atom{atom.NewAtomString("hello")},
		},
		{
			"int 32",
			"32",
			[]*atom.Atom{atom.NewAtomI(32)},
		},
		{
			"float 32",
			"32.0",
			[]*atom.Atom{atom.NewAtomF(32)},
		},
		{
			"true",
			"t",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"nil",
			"NIL",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		// arithmetic op
		{
			"add int int",
			"(+ 1 2)",
			[]*atom.Atom{atom.NewAtomI(3)},
		},
		{
			"add float.0 int",
			"(+ 1.0 1)",
			[]*atom.Atom{atom.NewAtomF(2)},
		},
		{
			"add float.1 int",
			"(+ 1.1 1)",
			[]*atom.Atom{atom.NewAtomF(2.1)},
		},
		{
			"add float float",
			"(+ 1.5 3.52)",
			[]*atom.Atom{atom.NewAtomF(5.02)},
		},
		{
			"add int int int",
			"(+ 1 2 3.0)",
			[]*atom.Atom{atom.NewAtomF(6)},
		},
		{
			"add add int",
			"(+ 1 (+ 2 5) 3 6)",
			[]*atom.Atom{atom.NewAtomI(17)},
		},
		{
			"add add float int",
			"(+ 1.2 (+ 1.2 (+ 2 3)))",
			[]*atom.Atom{atom.NewAtomF(7.4)},
		},
		{
			"add add add",
			"(+ (+ 1 1) (+ 2 2) (+ 3 3))",
			[]*atom.Atom{atom.NewAtomI(12)},
		},
		{
			"sub int int",
			"(- 3 2)",
			[]*atom.Atom{atom.NewAtomI(1)},
		},
		{
			"sub int float",
			"(- 3 2.0)",
			[]*atom.Atom{atom.NewAtomF(1)},
		},
		{
			"mul int int",
			"(* 5 2)",
			[]*atom.Atom{atom.NewAtomI(10)},
		},
		{
			"mul int float",
			"(* 5 2.0)",
			[]*atom.Atom{atom.NewAtomF(10)},
		},
		{
			"div int int",
			"(/ 100 10)",
			[]*atom.Atom{atom.NewAtomI(10)},
		},
		{
			"div int float",
			"(/ 100 10.0)",
			[]*atom.Atom{atom.NewAtomF(10)},
		},
		// logical op
		{
			"eq int int",
			"(= 1 1)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"eq int float",
			"(= 2 2.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"eq float float",
			"(= 4.4 4.4)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"!eq int int",
			"(= 1 99)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"!eq int float",
			"(= 99 9000.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"!eq float float",
			"(= 5.0 6.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"eq int int",
			"(= 1 1)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"ne int float",
			"(/= 2 2.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"ne float float",
			"(/= 4.4 4.4)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"!ne int int",
			"(/= 1 99)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"!ne int float",
			"(/= 99 9000.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"!ne float float",
			"(/= 5.0 6.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},

		{
			"lt int int",
			"(< 2 5)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"lt int int",
			"(< 20 5)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"lt int float",
			"(< 2 5.2)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"lt int float",
			"(< 20 5.99)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"lt float float",
			"(< 2.1 5.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"lt float float",
			"(< 20.0 5.900909)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"lt same int int",
			"(< 2 2)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"le int int",
			"(<= 2 5)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"le int int",
			"(<= 20 5)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"le int float",
			"(<= 2 5.2)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"le int float",
			"(<= 20 5.99)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"le float float",
			"(<= 2.1 5.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"le float float",
			"(<= 20.0 5.900909)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"le same int int",
			"(<= 2 2)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"gt int int",
			"(> 3 2)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"gt int float",
			"(> 3 2.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"gt float float",
			"(> 3.4 2.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"gt int int",
			"(> 3 20)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"gt int float",
			"(> 3 23.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"gt float float",
			"(> 3.4 2545.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"ge int int",
			"(>= 3 2)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"ge int float",
			"(>= 3 2.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"ge float float",
			"(>= 3.4 2.0)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"ge float int",
			"(>= 3.0 3)",
			[]*atom.Atom{atom.NewAtomTrue()},
		},
		{
			"ge int int",
			"(>= 3 20)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"ge int float",
			"(>= 3 21.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"ge float float",
			"(>= 3.4 23.0)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
		{
			"ge float int",
			"(>= 3.0 332323232)",
			[]*atom.Atom{atom.NewAtomNil()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			nodes, err := parse.Parse(tok)
			if err != nil {
				t.Fatal(err)
			}
			for i, node := range nodes {
				v, err := exec(node)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expect[i], v)
			}
		})
	}
}
