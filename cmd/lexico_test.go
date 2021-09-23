package lex_analyser

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
)

func createFile() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		//"/test/file.txt",
		"file1.txt",
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

func TestGetFilename(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)
}
