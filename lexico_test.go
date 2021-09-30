package lex_analyser

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

func TestIsIdentifierStarter(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    string
		expectedOuput bool
	}{
		{
			name:          "isIdentifierStarterAlphaLowCase",
			inputValue:    "a",
			expectedOuput: true,
		},
		{
			name:          "isIdentifierStarterAlphaUpperCase",
			inputValue:    "M",
			expectedOuput: true,
		},
		{
			name:          "isIdentifierStarterUnderscore",
			inputValue:    "_",
			expectedOuput: true,
		},
		{
			name:          "isIdentifierStarterAlphaNumber",
			inputValue:    "0",
			expectedOuput: false,
		},
		{
			name:          "isIdentifierStarterAlphaSymbol",
			inputValue:    "!",
			expectedOuput: false,
		},
	}

	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			input := bufio.NewReader(strings.NewReader(subtest.inputValue))
			lex := NewLexico(input)
			lex.yytext[0], _ = lex.getRune()
			if result := lex.isIdentifierStarter(lex.yytext[0]); result != subtest.expectedOuput {
				t.Errorf("Test %s, value expected '%t' and get '%t'\n", subtest.name, subtest.expectedOuput, result)
			}
		})
	}

}

func createBuffer(size int, bufferContent string) *Buffer {
	b := NewBuffer(size)
	if bufferContent != "" {
		for _, c := range bufferContent {
			b.putRune(c)
		}
	}
	return b
}

func TestGetRune(t *testing.T) {

	tests := []struct {
		name           string
		input          string
		bufferContent  string
		expectedResult rune
		expectedErr    error
	}{
		{
			name:           "BufferIsNullOK",
			input:          "A",
			bufferContent:  "",
			expectedResult: 'A',
			expectedErr:    nil,
		},
		{
			name:           "BufferIsNotNullOK",
			input:          "A",
			bufferContent:  "B",
			expectedResult: 'B',
			expectedErr:    nil,
		},
		{
			name:           "BufferIsNullErr",
			input:          "",
			bufferContent:  "",
			expectedResult: 0,
			expectedErr:    errors.New("EOF"),
		},
	}

	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			lex := NewLexico(strings.NewReader(subtest.input))
			lex.buffer = createBuffer(len(subtest.bufferContent)+1 /*Don't want nil buffers*/, subtest.bufferContent)
			result, err := lex.getRune()
			if result != subtest.expectedResult || err != nil && (err.Error() != subtest.expectedErr.Error()) {
				t.Errorf("Test %s, expected value '%s'and error '%s' and get value '%s' and error '%s\n", subtest.name,
					string(subtest.expectedResult), string(subtest.expectedErr.Error()), string(result), string(err.Error()))
			}

		})
	}

}

func TestPutCharBack(t *testing.T) {
	tests := []struct {
		name          string
		bufferContent string
		bufferSize    int
		expectedError error
	}{
		{
			name:          "bufferIsNotFull",
			bufferContent: "Hi",
			bufferSize:    len("Hi") + 1,
			expectedError: nil,
		},
		{
			name:          "bufferOverflow",
			bufferContent: "Hi",
			bufferSize:    len("Hi"),
			expectedError: ErrorBufferOverflow,
		},
	}

	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			lex := NewLexico(strings.NewReader("Test"))
			lex.buffer = createBuffer(subtest.bufferSize, subtest.bufferContent)
			err := lex.putCharBack('A')
			if err != subtest.expectedError {
				t.Errorf("Test %s, expected error '%s' and get error '%s\n", subtest.name, subtest.expectedError.Error(), err.Error())
			}
		})
	}
}
