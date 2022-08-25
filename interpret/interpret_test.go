package interpret

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/parse"
	"github.com/x0y14/goclisp/tokenize"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*Atom
	}{
		{
			"string",
			"\"hello\"",
			[]*Atom{NewAtomString("hello")},
		},
		{
			"int 32",
			"32",
			[]*Atom{NewAtomI(32)},
		},
		{
			"float 32",
			"32.0",
			[]*Atom{NewAtomF(32)},
		},
		{
			"true",
			"t",
			[]*Atom{NewAtomTrue()},
		},
		{
			"nil",
			"NIL",
			[]*Atom{NewAtomNil()},
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
				atom, err := exec(node)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expect[i], atom)
			}
		})
	}
}
