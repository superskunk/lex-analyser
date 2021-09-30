package lex_analyser

type TokenType int64
type token []rune

const (
	CurlyBracesOpenedToken TokenType = iota
	CurlyBracesClosedToken
	ParenthesisOpenedToken
	ParenthesisClosedToken
	QuoteToken
	DoubleQuoteToken
	AmpersandToken
	PipeToken
	EqualToken
	PeriodToken
	CommaToken
	ColonToken
	SemiColonToken
	AddToken
	DashToken
	QuestionMarkOpenedToken
	QuestionMarkClosedToken
	ExclamationMarkOpenedToken
	ExclamationMarkClosedToken
	XORBitwiseToken
	GreaterToken
	LessThanToken
	PercentageToken
	AsteriskToken
	SlashToken
	BackSlashToken
	UnderscoreToken
	CRToken
	BlankToken
	TabToken
	// From here tokens represented are not y symbols arrayd
	IdentifierToken
	IntegerConstantToken
	FloatConstantToken
	CharConstantToken
	StringConstantToken
	EOFToken     = -100
	UnknownToken = -101
)

// blanks is a uint that has four 1 bits activated (e.g: 1 << 10 --> 2^10 = 1024 = (64 bits) 0000 ... 0100 0000 0000)
const blanks = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '

var symbols = [...]rune{'{', '}', '(', ')', '\'', '"', '&', '|', '=', '.', ',', ':', ';', '+', '-', '¿', '?', '¡', '!', '^', '>', '<', '%', '*', '/', '\\', '_', '\n', ' ', '\t'}

//var blanks = [...]rune{'\n', ' ', '\t'}
var tokenNames = [...]string{
	"CurlyBracesOpenedToken",
	"CurlyBracesClosedToken",
	"ParenthesisOpenedToken",
	"ParenthesisClosedToken",
	"QuoteToken",
	"DoubleQuoteToken",
	"AmpersandToken",
	"PipeToken",
	"EqualToken",
	"PeriodToken",
	"CommaToken",
	"ColonToken",
	"SemiColonToken",
	"AddToken",
	"DashToken",
	"QuestionMarkOpenedToken",
	"QuestionMarkClosedToken",
	"ExclamationMarkOpenedToken",
	"ExclamationMarkClosedToken",
	"XORBitwiseToken",
	"GreaterToken",
	"LessThanToken",
	"PercentageToken",
	"AsteriskToken",
	"SlashToken",
	"BackSlashToken",
	"UnderscoreToken",
	"CRToken",
	"BlankToken",
	"TabToken",
	"IdentifierToken",
	"IntegerConstantToken",
	"FloatConstantToken",
	"CharConstantToken",
	"StringConstantToken",
	"EOFToken",
	"UnknownToken",
}

func (t TokenType) Index() int {
	return int(t)
}

func (t TokenType) Value() rune {
	return symbols[t]
	//return token(fmt.Sprintf("%v", symbols[t-1]))
}

func (t TokenType) Name() string {
	switch t {
	case EOFToken:
		return "EOFToken"
	case UnknownToken:
		return "UnknownToken"
	}
	return tokenNames[t]
}

func getTokenType(r rune) TokenType {
	for k, v := range symbols {
		if v == r {
			return TokenType(k)
		}
	}
	return UnknownToken
}

func isBlank(r rune) bool {
	return blanks&(1<<uint(r)) != 0
}
