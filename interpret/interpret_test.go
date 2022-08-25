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
