package lex_analyser

import (
	"io"
	"unicode"
)

func (l *lexico) analizeNumber() (token, TokenType, error) {
	var nToken token
	var err error

	for {
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getRune(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
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
			if l.yytext[l.yytextPointer], err = l.getRune(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
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
		if l.yytext[l.yytextPointer], err = l.getRune(); err == nil && l.yytext[l.yytextPointer] != AddToken.Value() &&
			l.yytext[l.yytextPointer] != DashToken.Value() && !unicode.IsDigit(l.yytext[l.yytextPointer]) {
			nToken := l.yytext[:l.yytextPointer]
			return nToken, UnknownToken, ErrorWrongFloatConst
		}

		for {
			l.yytextPointer++
			if l.yytext[l.yytextPointer], err = l.getRune(); err == nil && unicode.IsDigit(l.yytext[l.yytextPointer]) {
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
