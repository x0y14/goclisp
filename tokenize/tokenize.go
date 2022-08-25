package tokenize

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

var userInput []rune
var reservedSymbols []rune
var reservedCompositeSymbols []string

// position
var l int
var lp int
var wp int

func init() {
	reservedSymbols = []rune{
		'{', '}',
		'(', ')',
		'[', ']',
		'+', '-', '*', '/',
		'=', '>', '<',
		'&',
		'#',
		'\'',
		';',
	}
	reservedCompositeSymbols = []string{
		"/=", ">=", "<=",
	}
}

func startWith(q string) bool {
	qRunes := []rune(q)
	for i := 0; i < len(qRunes); i++ {
		if userInput[wp+i] != qRunes[i] {
			return false
		}
	}
	return true
}

func isIdentRune(r rune) bool {
	return ('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9') ||
		('_' == r)
}

func consumeIdent() string {
	var ident string
	for wp < len(userInput) {
		if isIdentRune(userInput[wp]) {
			ident += string(userInput[wp])
			lp++
			wp++
		} else {
			break
		}
	}
	return ident
}

func consumeString() string {
	var str string
	// "
	str += "\""
	lp++
	wp++

	for '"' != userInput[wp] {
		// escaped double quotation
		if '\\' == userInput[wp] && '"' == userInput[wp+1] {
			str += "\\\""
			lp += 2
			wp += 2
			continue
		}

		str += string(userInput[wp])
		lp++
		wp++
	}

	// "
	str += "\""
	lp++
	wp++
	return str
}

func consumeNumber() string {
	var numStr string
	for wp < len(userInput) &&
		(unicode.IsDigit(userInput[wp]) ||
			'.' == userInput[wp]) {
		numStr += string(userInput[wp])
		lp++
		wp++
	}
	return numStr
}

func Tokenize(in string) *Token {
	// initialize
	userInput = []rune(in)
	l = 1
	lp = 0
	wp = 0
	var head Token
	cur := &head

userInputLoop:
	for wp < len(userInput) {
		// White
		// whitespace
		if ' ' == userInput[wp] || '\t' == userInput[wp] {
			lp++
			wp++
			continue
		}
		// newline
		if '\n' == userInput[wp] {
			l++
			lp = 0
			wp++
			continue
		}

		// Reserved
		// composite symbols
		for _, comp := range reservedCompositeSymbols {
			if startWith(comp) {
				cur = NewTokenReserved(cur, NewPosition(l, lp, wp), comp)
				lp += len(comp)
				wp += len(comp)
				continue userInputLoop
			}
		}
		// single symbols
		if strings.ContainsRune(string(reservedSymbols), userInput[wp]) {
			cur = NewTokenReserved(cur, NewPosition(l, lp, wp), string(userInput[wp]))
			lp++
			wp++
			continue
		}

		// Ident
		if !unicode.IsDigit(userInput[wp]) && isIdentRune(userInput[wp]) {
			cur = NewTokenIdent(cur, NewPosition(l, lp, wp), consumeIdent())
			// lp, wp had increased by consumeIdent.
			continue
		}

		// String
		if '"' == userInput[wp] {
			cur = NewTokenString(cur, NewPosition(l, lp, wp), consumeString())
			// lp, wp had increased by consumeString.
			continue
		}

		// Number
		if unicode.IsDigit(userInput[wp]) {
			pos := NewPosition(l, lp, wp)
			numStr := consumeNumber()
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				log.Fatalf("[%d:%d] failed to parse number: %s", pos.LineNo, pos.LpBegin, numStr)
			}
			cur = NewTokenNumber(cur, pos, num, numStr)
			// lp, wp had increased by consumeNumber
			continue
		}

		log.Fatalf("[%d:%d] unexpected charactor: %s", l, lp, string(userInput[wp]))
	}

	NewTokenEof(cur, NewPosition(l, lp, wp))
	return head.Next
}
