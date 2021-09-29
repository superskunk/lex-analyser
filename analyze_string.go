package lex_analyser

import "io"

func (l *lexico) analyzeString() (token, TokenType, error) {
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
