package main

import (
	"fmt"
	"io"
	"log"

	"github.com-superskunk/superskunk/lex_analyser"
)

func main() {
	fmt.Println("Starting analisis")

	lex := lex_analyser.NewLexico("file.txt")

	err := lex.Open()
	if err != nil {
		log.Fatalf("Unable to open file %s", lex.FileName)
	}
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
