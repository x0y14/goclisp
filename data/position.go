package data

type Position struct {
	LineNo  int // 1 <= n
	LpBegin int // LinePosition Position in Line
	WpBegin int // WholePosition Position in InputText
}

func NewPosition(line, lp, wp int) *Position {
	return &Position{
		LineNo:  line,
		LpBegin: lp,
		WpBegin: wp,
	}
}
