package lex_analyser

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func createFile() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"/test/file.txt",
		//os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func removeFile() {

}

func Test_Open(t *testing.T) {
	createFile()
	fmt.Print("hola")
}
