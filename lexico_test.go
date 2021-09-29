package lex_analyser

import (
	"bufio"
	"strings"
	"testing"
)

func TestIsIdentifierStarter(t *testing.T) {
	//scanner := bufio.NewScanner(strings.NewReader(validIdentifier))
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
