package lex_analyser

import (
	"io"
	"unicode"
)

func (l *lexico) isIdentifierStarter(c rune) bool {
	return c == UnderscoreToken.Value() || unicode.IsLetter(c)
}

func (l *lexico) isIdentifierChar(c rune) bool {
	return c == UnderscoreToken.Value() || unicode.IsLetter(c) || unicode.IsDigit(c)
}

func (l *lexico) analizeIdentifier() (token, TokenType, error) {
	var nToken token
	var err error

	for {
		l.yytextPointer++
		if l.yytext[l.yytextPointer], err = l.getChar(); err == nil && l.isIdentifierChar(l.yytext[l.yytextPointer]) {
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
