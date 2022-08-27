package tokenize

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/goclisp/data"
	"testing"
)

func NewL1Position(lp int) *data.Position {
	return data.NewPosition(1, lp, lp)
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *data.Token
	}{
		{
			"wild plus",
			"1+1",
			&data.Token{
				Kind:     data.TkInt,
				Position: data.NewPosition(1, 0, 0),
				Num:      1,
				Str:      "1",
				Next: &data.Token{
					Kind:     data.TkReserved,
					Position: data.NewPosition(1, 1, 1),
					Num:      0,
					Str:      "+",
					Next: &data.Token{
						Kind:     data.TkInt,
						Position: data.NewPosition(1, 2, 2),
						Num:      1,
						Str:      "1",
						Next: &data.Token{
							Kind:     data.TkEof,
							Position: data.NewPosition(1, 3, 3),
							Num:      0,
							Str:      "",
							Next:     nil,
						},
					},
				},
			},
		},
		{
			"wild string",
			"\"string\"",
			&data.Token{
				Kind:     data.TkString,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "string",
				Next: &data.Token{
					Kind:     data.TkEof,
					Position: NewL1Position(len("\"string\"")),
					Num:      0,
					Str:      "",
					Next:     nil,
				},
			},
		},
		{
			"wild comp symbol",
			"/=",
			&data.Token{
				Kind:     data.TkReserved,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "/=",
				Next: &data.Token{
					Kind:     data.TkEof,
					Position: NewL1Position(len("/=")),
					Num:      0,
					Str:      "",
					Next:     nil,
				},
			},
		},
		{
			"wild single symbol",
			"+",
			&data.Token{
				Kind:     data.TkReserved,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "+",
				Next: &data.Token{
					Kind:     data.TkEof,
					Position: NewL1Position(1),
					Num:      0,
					Str:      "",
					Next:     nil,
				},
			},
		},
		{
			"wild ident",
			"ident",
			&data.Token{
				Kind:     data.TkIdent,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "ident",
				Next: &data.Token{
					Kind:     data.TkEof,
					Position: NewL1Position(len("ident")),
					Num:      0,
					Str:      "",
					Next:     nil,
				},
			},
		},
		{
			"plus",
			"(+ 2.2 -30.0)",
			&data.Token{
				Kind:     data.TkReserved,
				Position: NewL1Position(len("")),
				Num:      0,
				Str:      "(",
				Next: &data.Token{
					Kind:     data.TkReserved,
					Position: NewL1Position(len("(")),
					Num:      0,
					Str:      "+",
					Next: &data.Token{
						Kind:     data.TkFloat,
						Position: NewL1Position(len("(+ ")),
						Num:      2.2,
						Str:      "2.2",
						Next: &data.Token{
							Kind:     data.TkReserved,
							Position: NewL1Position(len("(+ 2.2 ")),
							Num:      0,
							Str:      "-",
							Next: &data.Token{
								Kind:     data.TkFloat,
								Position: NewL1Position(len("(+ 2.2 -")),
								Num:      30,
								Str:      "30.0",
								Next: &data.Token{
									Kind:     data.TkReserved,
									Position: NewL1Position(len("(+ 2.2 -30.0")),
									Num:      0,
									Str:      ")",
									Next: &data.Token{
										Kind:     data.TkEof,
										Position: NewL1Position(len("(+ 2.2 -30.0)")),
										Num:      0,
										Str:      "",
										Next:     nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, tok)
		})
	}
}
