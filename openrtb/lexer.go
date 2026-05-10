package openrtb

import (
	"fmt"
)

type tokenKind int8

const (
	tokString tokenKind = iota
	tokNumber
	tokTrue
	tokFalse
	tokNull
	tokLBrace
	tokRBrace
	tokLBracket
	tokRBracket
	tokColon
	tokComma
	tokEOF
)

type token struct {
	kind tokenKind
	val  []byte
}

type lexer struct {
	input []byte
	pos   int
}

func newLexer(data []byte) *lexer {
	return &lexer{input: data}
}

func (l *lexer) next() (token, error) {
	for l.pos < len(l.input) && (l.input[l.pos] == ' ' || l.input[l.pos] == '\n' ||
		l.input[l.pos] == '\t' || l.input[l.pos] == '\r') {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return token{kind: tokEOF}, nil
	}

	ch := l.input[l.pos]

	switch ch {
	case '{':
		l.pos++
		return token{kind: tokLBrace}, nil
	case '}':
		l.pos++
		return token{kind: tokRBrace}, nil
	case '[':
		l.pos++
		return token{kind: tokLBracket}, nil
	case ']':
		l.pos++
		return token{kind: tokRBracket}, nil
	case ':':
		l.pos++
		return token{kind: tokColon}, nil
	case ',':
		l.pos++
		return token{kind: tokComma}, nil
	case '"':
		return l.scanString()
	case 't':
		tokLength := len("true")
		if l.pos+tokLength <= len(l.input) && string(l.input[l.pos:l.pos+tokLength]) == "true" {
			l.pos += tokLength
			return token{kind: tokTrue}, nil
		}
	case 'f':
		tokLength := len("false")
		if l.pos+tokLength <= len(l.input) && string(l.input[l.pos:l.pos+tokLength]) == "false" {
			l.pos += tokLength
			return token{kind: tokFalse}, nil
		}
	case 'n':
		tokLength := len("null")
		if l.pos+tokLength <= len(l.input) && string(l.input[l.pos:l.pos+tokLength]) == "null" {
			l.pos += tokLength
			return token{kind: tokNull}, nil
		}
	default:
		if ch == '-' || (ch >= '0' && ch <= '9') {
			start := l.pos
			for l.pos < len(l.input) && ((l.input[l.pos] >= '0' && l.input[l.pos] <= '9') ||
				l.input[l.pos] == '.' || l.input[l.pos] == 'e' || l.input[l.pos] == 'E' ||
				l.input[l.pos] == '+' || l.input[l.pos] == '-') {
				l.pos++
			}
			return token{kind: tokNumber, val: l.input[start:l.pos]}, nil
		}
	}

	return token{}, fmt.Errorf("lexer: unexpected byte %q at pos %v", ch, l.pos)
}

func (l *lexer) scanString() (token, error) {
	l.pos++
	var buf []byte

	for l.pos < len(l.input) {
		ch := l.input[l.pos]

		if ch == '"' {
			l.pos++
			return token{kind: tokString, val: buf}, nil
		}

		if ch == '\\' {
			l.pos++
			if l.pos >= len(l.input) {
				return token{}, fmt.Errorf("lexer: unterminated escape")
			}
			switch l.input[l.pos] {
			case '"':
				buf = append(buf, '"')
			case '\\':
				buf = append(buf, '\\')
			case '/':
				buf = append(buf, '/')
			case 'n':
				buf = append(buf, '\n')
			case 'r':
				buf = append(buf, '\r')
			case 't':
				buf = append(buf, '\t')
			case 'b':
				buf = append(buf, '\b')
			case 'f':
				buf = append(buf, '\f')
			default:
				return token{}, fmt.Errorf("lexer: unsupported escape sequence %q", l.input[l.pos])
			}
			l.pos++
			continue
		}

		buf = append(buf, ch)
		l.pos++
	}

	return token{}, fmt.Errorf("lexer: unterminated string")
}

func (l *lexer) scanRaw() ([]byte, error) {
	return nil, nil
}
