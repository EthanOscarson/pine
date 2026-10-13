package compiler

type Scanner struct {
	source string
	start int
	current int
	line int
}

func newScanner(source string) Scanner {
	s := Scanner{
		source: string,
		start: 0,
		current: 0,
		line: 1,
	}
	return s;
}

func (s *Scanner) scanToken () {
	s.start = s.current
	if (isAtEnd()) {
		return makeToken(TOKEN_EOF)
	}
	return errorToken("Unexpected character.")
}

func (s *Scanner) isAtEnd () bool {
	return s.current == len(s.source)
}
