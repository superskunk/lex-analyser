package lex_analyser

import (
	"io"
	"log"
	"os"
	"unicode"
)

const TokenMaxLong = 100

type LexAnalizer interface {
	isIdentifierStarter(c rune) bool
	isIdentifierChar(c rune) bool
}

type lexico struct {
	fileName      string
	file          *os.File
	buffer        *Buffer
	yytext        token
	yytextPointer int
	yylineno      int
	analizer      LexAnalizer
}

func NewLexico(fileName string) *lexico {
	lex := &lexico{}
	lex.fileName = fileName
	lex.buffer = NewBuffer(BufferMaxSize)
	lex.yytext = make(token, TokenMaxLong)
	lex.analizer = NewAnalizer()
	return lex
}

func (l *lexico) Open() {
	f, err := os.Open(l.fileName)
	l.file = f
	if err != nil {
		log.Fatal(err)
	}
}

func (l *lexico) Close() {
	l.file.Close()
}

func (l *lexico) readByteFromFile() (rune, error) {
	var buf [1]byte
	_, err := l.file.Read(buf[:])
	return (rune)(buf[0]), err
}

func (l *lexico) getChar() (rune, error) {

	if c, err := l.buffer.getRune(); err == nil {
		return c, nil
	}
	return l.readByteFromFile()
}

func (l *lexico) putCharBack(c rune) {
	if err := l.buffer.putRune(c); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func (l *lexico) skipBlanks() error {
	var err error
	for {
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && isBlank(l.yytext[l.yytextPointer]) {
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

func (l *lexico) analizeIdentifier() (token, TokenType, error) {
	var nToken token
	var err error

	for {
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && l.analizer.isIdentifierChar(l.yytext[l.yytextPointer]) {
			continue
		}
		break
	}

	nToken = l.yytext[:l.yytextPointer]
	l.putCharBack(l.yytext[l.yytextPointer])
	tokenType := IdentifierToken

	if err == io.EOF {
		tokenType = EOFToken
	}

	return nToken, tokenType, err

}

func (l *lexico) analizeNumber() (token, TokenType, error) {
	var nToken token
	var err error

	for {
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
			continue
		}
		break
	}

	if err == io.EOF {
		return l.yytext, EOFToken, err
	}

	// We replace E by .0E
	// e.g: 50E = 50.0E
	if l.yytext[l.yytextPointer] == 'E' {
		exp := l.yytext[l.yytextPointer : l.yytextPointer+3]
		// as exp is a subslice from l.yytext, changes made in exp have effect in l.yytext
		copy(exp, []rune{PeriodToken.Value(), '0', 'E'})
		l.yytextPointer += 2

	}

	// If it is not a '.' neither an 'E', then it is an int

	if l.yytext[l.yytextPointer] != PeriodToken.Value() && l.yytext[l.yytextPointer] != 'E' {
		nToken = l.yytext[:l.yytextPointer]
		l.putCharBack(l.yytext[l.yytextPointer])
		//l.buffer.putRune(l.yytext[l.yytextPointer])
		return nToken, IntegerConstantToken, nil
	}
	return nToken, FloatConstantToken, nil
}

func (l *lexico) analizeFloatNumber() (token, TokenType, error) {
	var nToken token
	var err error

	if l.yytext[l.yytextPointer] == PeriodToken.Value() {
		// Read while digits
		for {
			l.yytextPointer++
			if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
				continue
			}
			break
		}

		if err == io.EOF {
			return l.yytext, EOFToken, err
		}
	}

	if l.yytext[l.yytextPointer] == 'E' && l.yytext[l.yytextPointer-1] == PeriodToken.Value() {
		exp := l.yytext[l.yytextPointer : l.yytextPointer+2]
		copy(exp, []rune{'0', 'E'})
		l.yytextPointer += 2
	}

	if l.yytext[l.yytextPointer] == 'E' {
		// Now we can read a '+', '-' or a digit
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && l.yytext[l.yytextPointer] != AddToken.Value() &&
			l.yytext[l.yytextPointer] != DashToken.Value() && !unicode.IsDigit(l.yytext[l.yytextPointer]) {
			nToken := l.yytext[:l.yytextPointer]
			return nToken, UnknownToken, ErrorWrongFloatConst
		}

		for {
			l.yytextPointer++
			if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
				continue
			}
			break
		}

		if err == io.EOF {
			return l.yytext, EOFToken, err
		}

	}

	nToken = l.yytext[:l.yytextPointer]
	l.putCharBack(l.yytext[l.yytextPointer])
	if l.yytextPointer == 1 {
		return nToken, PeriodToken, nil
	}
	return nToken, FloatConstantToken, nil

}

func (l *lexico) analizeString() (token, TokenType, error) {
	var nToken token
	var err error

	quoteType := l.yytext[l.yytextPointer]

	l.yytextPointer++
	if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && l.yytext[l.yytextPointer] == BackSlashToken.Value() {
		l.yytextPointer++
		l.yytext[l.yytextPointer], _ = l.getChar()
	}

	if err == io.EOF {
		return nToken, EOFToken, err
	}

	for {
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil &&
			l.yytext[l.yytextPointer] != QuoteToken.Value() &&
			l.yytext[l.yytextPointer] != DoubleQuoteToken.Value() {
			continue
		}
		break
	}

	if err == io.EOF {
		return nToken, EOFToken, err
	}

	nToken = l.yytext[:l.yytextPointer+1]

	if nToken[l.yytextPointer] != quoteType {
		return nToken, UnknownToken, ErrorWrongStringConst
	}

	return nToken, StringConstantToken, nil
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

	if l.analizer.isIdentifierStarter(l.yytext[l.yytextPointer]) {
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
		return l.analizeString()
	}

	nToken = l.yytext[:l.yytextPointer+1]

	return nToken, getTokenType(nToken[0]), nil

}
