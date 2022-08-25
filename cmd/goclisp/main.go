package main

import (
	"bufio"
	"fmt"
	"github.com/x0y14/goclisp/interpret"
	"github.com/x0y14/goclisp/parse"
	"github.com/x0y14/goclisp/tokenize"
	"os"
)

func main() {
	userInput := ""
	scanner := bufio.NewScanner(os.Stdin)
	for {
		userInput = ""
		fmt.Print("* ")
		scanner.Scan()
		line := scanner.Text()
		if len(line) != 0 {
			userInput += line
		}

		if len(userInput) > 0 {
			// try tokenize & parse
			token, err := tokenize.Tokenize(userInput)
			if err != nil {
				continue
			}
			nodes, err := parse.Parse(token)
			if err != nil {
				continue
			}
			// do it!!
			err = interpret.Interpret(nodes)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
