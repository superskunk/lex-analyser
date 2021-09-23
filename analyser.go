package lex_analyser

import (
	"errors"
	"unicode"
)

var (
	ErrorWrongFloatConst  error = errors.New("error, float const is not well formed")
	ErrorWrongCharConst   error = errors.New("error, char const is not well formed")
	ErrorWrongStringConst error = errors.New("error, string const is not well formed")
)

type analizer struct {
}

func NewAnalizer() LexAnalizer {
	return &analizer{}
}

func (a *analizer) isIdentifierStarter(c rune) bool {
	return c == UnderscoreToken.Value() || unicode.IsLetter(c)
}

func (a *analizer) isIdentifierChar(c rune) bool {
	return c == UnderscoreToken.Value() || unicode.IsLetter(c) || unicode.IsDigit(c)
}
