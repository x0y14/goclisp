package interpret

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/data"
	"github.com/x0y14/goclisp/parse"
	"github.com/x0y14/goclisp/tokenize"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*data.Data
	}{
		// atom
		{
			"string",
			"\"hello\"",
			[]*data.Data{data.NewDataString("hello")},
		},
		{
			"int 32",
			"32",
			[]*data.Data{data.NewDataInt(32)},
		},
		{
			"float 32",
			"32.0",
			[]*data.Data{data.NewDataFloat(32)},
		},
		{
			"true",
			"t",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"nil",
			"NIL",
			[]*data.Data{data.NewDataNil()},
		},
		// arithmetic op
		{
			"add int int",
			"(+ 1 2)",
			[]*data.Data{data.NewDataInt(3)},
		},
		{
			"add float.0 int",
			"(+ 1.0 1)",
			[]*data.Data{data.NewDataFloat(2)},
		},
		{
			"add float.1 int",
			"(+ 1.1 1)",
			[]*data.Data{data.NewDataFloat(2.1)},
		},
		{
			"add float float",
			"(+ 1.5 3.52)",
			[]*data.Data{data.NewDataFloat(5.02)},
		},
		{
			"add int int int",
			"(+ 1 2 3.0)",
			[]*data.Data{data.NewDataFloat(6)},
		},
		{
			"add add int",
			"(+ 1 (+ 2 5) 3 6)",
			[]*data.Data{data.NewDataInt(17)},
		},
		{
			"add add float int",
			"(+ 1.2 (+ 1.2 (+ 2 3)))",
			[]*data.Data{data.NewDataFloat(7.4)},
		},
		{
			"add add add",
			"(+ (+ 1 1) (+ 2 2) (+ 3 3))",
			[]*data.Data{data.NewDataInt(12)},
		},
		{
			"sub int int",
			"(- 3 2)",
			[]*data.Data{data.NewDataInt(1)},
		},
		{
			"sub int float",
			"(- 3 2.0)",
			[]*data.Data{data.NewDataFloat(1)},
		},
		{
			"mul int int",
			"(* 5 2)",
			[]*data.Data{data.NewDataInt(10)},
		},
		{
			"mul int float",
			"(* 5 2.0)",
			[]*data.Data{data.NewDataFloat(10)},
		},
		{
			"div int int",
			"(/ 100 10)",
			[]*data.Data{data.NewDataInt(10)},
		},
		{
			"div int float",
			"(/ 100 10.0)",
			[]*data.Data{data.NewDataFloat(10)},
		},
		// logical op
		{
			"eq int int",
			"(= 1 1)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"eq int float",
			"(= 2 2.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"eq float float",
			"(= 4.4 4.4)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"!eq int int",
			"(= 1 99)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"!eq int float",
			"(= 99 9000.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"!eq float float",
			"(= 5.0 6.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"eq int int",
			"(= 1 1)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"ne int float",
			"(/= 2 2.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"ne float float",
			"(/= 4.4 4.4)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"!ne int int",
			"(/= 1 99)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"!ne int float",
			"(/= 99 9000.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"!ne float float",
			"(/= 5.0 6.0)",
			[]*data.Data{data.NewDataTrue()},
		},

		{
			"lt int int",
			"(< 2 5)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"lt int int",
			"(< 20 5)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"lt int float",
			"(< 2 5.2)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"lt int float",
			"(< 20 5.99)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"lt float float",
			"(< 2.1 5.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"lt float float",
			"(< 20.0 5.900909)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"lt same int int",
			"(< 2 2)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"le int int",
			"(<= 2 5)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"le int int",
			"(<= 20 5)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"le int float",
			"(<= 2 5.2)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"le int float",
			"(<= 20 5.99)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"le float float",
			"(<= 2.1 5.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"le float float",
			"(<= 20.0 5.900909)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"le same int int",
			"(<= 2 2)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"gt int int",
			"(> 3 2)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"gt int float",
			"(> 3 2.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"gt float float",
			"(> 3.4 2.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"gt int int",
			"(> 3 20)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"gt int float",
			"(> 3 23.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"gt float float",
			"(> 3.4 2545.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"ge int int",
			"(>= 3 2)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"ge int float",
			"(>= 3 2.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"ge float float",
			"(>= 3.4 2.0)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"ge float int",
			"(>= 3.0 3)",
			[]*data.Data{data.NewDataTrue()},
		},
		{
			"ge int int",
			"(>= 3 20)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"ge int float",
			"(>= 3 21.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"ge float float",
			"(>= 3.4 23.0)",
			[]*data.Data{data.NewDataNil()},
		},
		{
			"ge float int",
			"(>= 3.0 332323232)",
			[]*data.Data{data.NewDataNil()},
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
				v, err := eval(data.GlobalStorage, node)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expect[i], v)
			}
		})
	}
}
