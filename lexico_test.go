package lex_analyser

import (
	"os"
	"testing"
)

func createContentFile() []byte {

	contentString := `Hello,
	this is an Identifier,
	this is an integer 200
	this is a float number 103.3
	this is another float number 28.3E+3
	this is a string "Good Morgning"
	`
	return []byte(contentString)
}

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("file_lexito_test.txt")
	if err != nil {
		return "", err
	}
	// write some data to f
	content := createContentFile()
	_, err = f.Write(content)

	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name             string
		createF          func(t *testing.T) (string, error)
		expectedFileName string
		errorOpeningFile bool
	}{
		{
			name:             "openFileOK",
			createF:          createFile,
			expectedFileName: "file_lexito_test.txt",
			errorOpeningFile: false,
		},
		{
			name: "openFileNOK",
			createF: func(t *testing.T) (string, error) {
				return "A_File_Does_Not_Exists.txt", nil
			},
			errorOpeningFile: true,
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			fileName, _ := subtest.createF(t)
			lex := NewLexico(fileName)
			lex.Open()
			switch subtest.errorOpeningFile {
			case true:
				if lex.file != nil {
					t.Errorf("Expected lex.file was nil, get lex.file non-nil\n")
				}
			case false:
				if lex.FileName != subtest.expectedFileName {
					t.Errorf("Expected the file %s to be opened, get the file %s opened\n", subtest.expectedFileName, lex.FileName)
				}
			}
		})
	}
}
