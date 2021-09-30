package lex_analyser

import (
	"errors"
	"io"
	"unicode"
)

var (
	ErrorWrongFloatConst  error = errors.New("error, float const is not well formed")
	ErrorWrongCharConst   error = errors.New("error, char const is not well formed")
	ErrorWrongStringConst error = errors.New("error, string const is not well formed")
)

const TokenMaxLong = 100

type LexAnalizer interface {
	isIdentifierStarter(c rune) bool
	isIdentifierChar(c rune) bool
}

type lexico struct {
	input         io.Reader
	buffer        *Buffer
	yytext        token
	yytextPointer int
	yylineno      int
}

func NewLexico(input io.Reader) *lexico {
	lex := &lexico{}
	lex.input = input
	lex.buffer = NewBuffer(BufferMaxSize)
	lex.yytext = make(token, TokenMaxLong)
	return lex
}

func (l *lexico) readByteFromInput() (rune, error) {
	var buf [1]byte
	_, err := l.input.Read(buf[:])
	a := (rune)(buf[0])
	return a, err
	//return (rune)(buf[0]), err
}

func (l *lexico) getRune() (rune, error) {

	if c, err := l.buffer.getRune(); err == nil {
		return c, nil
	}
	return l.readByteFromInput()
}

func (l *lexico) putCharBack(c rune) error {
	return l.buffer.putRune(c)

	// if err := l.buffer.putRune(c); err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }
}

func (l *lexico) skipBlanks() error {
	var err error
	for {
		if l.yytext[l.yytextPointer], err = l.getRune(); err == nil && isBlank(l.yytext[l.yytextPointer]) {
			if l.yytext[l.yytextPointer] == CRToken.Value() {
				l.yylineno++
			}
			continue
		}
		break
		// Check if it is the End of File
	}

	return err
}

func (l *lexico) NextToken() (token, TokenType, error) {
	//TODO
	var err error
	var nToken token
	var ttype TokenType

	l.yytextPointer = 0

	if err = l.skipBlanks(); err == io.EOF {
		return l.yytext, EOFToken, err
	}

	if l.isIdentifierStarter(l.yytext[l.yytextPointer]) {
		return l.analizeIdentifier()
	}

	// Check if it is a digit
	if unicode.IsDigit(l.yytext[l.yytextPointer]) {
		nToken, ttype, err = l.analizeNumber()

		if ttype == IntegerConstantToken {
			return nToken, ttype, err
		}
	}

	// In case is a "." or an "E" or it came from analizeNumber as a FloatNumber...
	if l.yytext[l.yytextPointer] == PeriodToken.Value() || l.yytext[l.yytextPointer] == 'E' {
		return l.analizeFloatNumber()
	}

	if l.yytext[l.yytextPointer] == QuoteToken.Value() || l.yytext[l.yytextPointer] == DoubleQuoteToken.Value() {
		return l.analyzeString()
	}

	nToken = l.yytext[:l.yytextPointer+1]

	return nToken, getTokenType(nToken[0]), nil

}
