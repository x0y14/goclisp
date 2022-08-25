package parse

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/tokenize"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*Node
	}{
		{
			"1+2",
			"(+ 1 2)",
			[]*Node{
				NewNodeWithArgs(Add, []*Node{
					NewNodeInt(1),
					NewNodeInt(2),
				}),
			},
		},
		{
			"1+2*3",
			"(+ 1 (* 2 3))",
			[]*Node{
				NewNodeWithArgs(Add, []*Node{
					NewNodeInt(1),
					NewNodeWithArgs(Mul, []*Node{
						NewNodeInt(2),
						NewNodeInt(3),
					}),
				}),
			},
		},
		{
			"hello world",
			"(format t \"hello, world\")",
			[]*Node{
				NewNodeWithArgs(Call, []*Node{
					NewNodeTrue(),
					NewNodeString("\"hello, world\""),
				}),
			},
		},
		{
			"add 4",
			"(+ 1 -2 3.3 4.0)",
			[]*Node{
				NewNodeWithArgs(Add, []*Node{
					NewNodeInt(1),
					NewNodeWithArgs(Sub, []*Node{
						NewNodeInt(0),
						NewNodeInt(2),
					}),
					NewNodeFloat(3.3),
					NewNodeFloat(4),
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			node, err := Parse(tok)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, node)
		})
	}
}
