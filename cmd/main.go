package main

import (
	"fmt"
	"io"

	"github.com/superskunk/analizer"
)

func main() {
	fmt.Println("Starting analisis")

	lex := analizer.NewLexico("file.txt")

	lex.Open()
	defer lex.Close()

	eof := false

	for !eof {
		token, typeToken, err := lex.NextToken()
		switch err {
		case io.EOF:
			eof = true
			fmt.Printf("EOF found\n")
		case nil:
			fmt.Printf("'%s'  -----------> '%s'\n", string(token), string(typeToken.Name()))
		default:
			fmt.Printf("'%s' ----------> '%s' -------> '%s'\n", string(token), string(typeToken.Name()), err.Error())
		}
	}
}
