package lexer

type TokenType uint8

const (
	Undefined TokenType = iota
	
	keywordBeg
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	keywordEnd
	
	literalStart
	IDENTIFIER
	STRING
	NUMBER
	literalEnd
	
	operatorStart
	PLUS
	MINUS
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	operatorEnd
	
	delimStart
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	delimEnd
)

type Token struct {
	tokenType TokenType
	lexme string
	
	
}