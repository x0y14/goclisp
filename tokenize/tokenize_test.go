package tokenize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewL1Position(lp int) *Position {
	return NewPosition(1, lp, lp)
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *Token
	}{
		{
			"wild plus",
			"1+1",
			&Token{
				Kind:     Number,
				Position: NewPosition(1, 0, 0),
				Num:      1,
				Str:      "1",
				Next: &Token{
					Kind:     Reserved,
					Position: NewPosition(1, 1, 1),
					Num:      0,
					Str:      "+",
					Next: &Token{
						Kind:     Number,
						Position: NewPosition(1, 2, 2),
						Num:      1,
						Str:      "1",
						Next: &Token{
							Kind:     Eof,
							Position: NewPosition(1, 3, 3),
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
			&Token{
				Kind:     String,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "\"string\"",
				Next: &Token{
					Kind:     Eof,
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
			&Token{
				Kind:     Reserved,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "/=",
				Next: &Token{
					Kind:     Eof,
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
			&Token{
				Kind:     Reserved,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "+",
				Next: &Token{
					Kind:     Eof,
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
			&Token{
				Kind:     Ident,
				Position: NewL1Position(0),
				Num:      0,
				Str:      "ident",
				Next: &Token{
					Kind:     Eof,
					Position: NewL1Position(len("ident")),
					Num:      0,
					Str:      "",
					Next:     nil,
				},
			},
		},
		{
			"plus",
			"(+ 2.2 -30..0)",
			&Token{
				Kind:     Reserved,
				Position: NewL1Position(len("")),
				Num:      0,
				Str:      "(",
				Next: &Token{
					Kind:     Reserved,
					Position: NewL1Position(len("(")),
					Num:      0,
					Str:      "+",
					Next: &Token{
						Kind:     Number,
						Position: NewL1Position(len("(+ ")),
						Num:      2.2,
						Str:      "2.2",
						Next: &Token{
							Kind:     Reserved,
							Position: NewL1Position(len("(+ 2.2 ")),
							Num:      0,
							Str:      "-",
							Next: &Token{
								Kind:     Number,
								Position: NewL1Position(len("(+ 2.2 -")),
								Num:      30,
								Str:      "30.0",
								Next: &Token{
									Kind:     Reserved,
									Position: NewL1Position(len("(+ 2.2 -30.0")),
									Num:      0,
									Str:      ")",
									Next: &Token{
										Kind:     Eof,
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
