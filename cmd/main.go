package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/superskunk/lex_analyser"
)

func main() {
	fmt.Println("Starting analisis")

	f, err := os.Open("file.txt")

	if err != nil {
		log.Fatalf("Unable to open file %s", "file.txt")
	}

	lex := lex_analyser.NewLexico(f)

	defer f.Close()

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
