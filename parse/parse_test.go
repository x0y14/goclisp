package parse

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/tokenize"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*data.Node
	}{
		{
			"1+2",
			"(+ 1 2)",
			[]*data.Node{
				data.NewNodeWithArgs(data.NdAdd, []*data.Node{
					data.NewNodeInt(1),
					data.NewNodeInt(2),
				}),
			},
		},
		{
			"1+2*3",
			"(+ 1 (* 2 3))",
			[]*data.Node{
				data.NewNodeWithArgs(data.NdAdd, []*data.Node{
					data.NewNodeInt(1),
					data.NewNodeWithArgs(data.NdMul, []*data.Node{
						data.NewNodeInt(2),
						data.NewNodeInt(3),
					}),
				}),
			},
		},
		{
			"hello world",
			"(format t \"hello, world\")",
			[]*data.Node{
				data.NewNodeCall("format", []*data.Node{
					data.NewNodeTrue(),
					data.NewNodeString("hello, world"),
				}),
			},
		},
		{
			"add 4",
			"(+ 1 -2 3.3 4.0)",
			[]*data.Node{
				data.NewNodeWithArgs(data.NdAdd, []*data.Node{
					data.NewNodeInt(1),
					data.NewNodeWithArgs(data.NdSub, []*data.Node{
						data.NewNodeInt(0),
						data.NewNodeInt(2),
					}),
					data.NewNodeFloat(3.3),
					data.NewNodeFloat(4),
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
