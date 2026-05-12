package openrtb

import (
	"fmt"
	"unsafe"
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
	input  []byte
	pos    int
	peeked *token
}

func newLexer(data []byte) *lexer {
	return &lexer{input: data}
}

func (l *lexer) peek() (token, error) {
	tok, err := l.next()
	if err != nil {
		return token{}, err
	}
	l.peeked = &tok

	return tok, nil
}

func (l *lexer) next() (token, error) {
	if l.peeked != nil {
		tok := *l.peeked
		l.peeked = nil
		return tok, nil
	}

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
	start := l.pos
	var buf []byte

	for l.pos < len(l.input) {
		ch := l.input[l.pos]

		if ch == '"' {
			var val []byte
			if buf == nil {
				val = l.input[start:l.pos]
			} else {
				val = buf
			}
			l.pos++
			return token{kind: tokString, val: val}, nil
		}

		if ch == '\\' {
			if buf == nil {
				buf = append(buf, l.input[start:l.pos]...)
			}
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

		if buf != nil {
			buf = append(buf, ch)
		}
		l.pos++
	}

	return token{}, fmt.Errorf("lexer: unterminated string")
}

func (l *lexer) scanRaw() ([]byte, error) {
	for l.pos < len(l.input) && (l.input[l.pos] == ' ' || l.input[l.pos] == '\n' ||
		l.input[l.pos] == '\t' || l.input[l.pos] == '\r') {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return nil, fmt.Errorf("scanRaw: unexpected end of input")
	}

	start := l.pos
	ch := l.input[l.pos]

	switch ch {
	case '"':
		l.pos++
		for l.pos < len(l.input) && l.input[l.pos] != '"' {
			if l.input[l.pos] == '\\' {
				l.pos += 2
			} else {
				l.pos++
			}
		}
		if l.pos >= len(l.input) {
			return nil, fmt.Errorf("scanRaw: unterminated string")
		}
		l.pos++
	case '{', '[':
		depth := 1
		inString := false
		l.pos++
		for depth != 0 {
			if l.pos >= len(l.input) {
				return nil, fmt.Errorf("scanRaw: unterminated object/array")
			}

			if inString {
				if l.input[l.pos] == '\\' {
					l.pos += 2
					continue
				}
				if l.input[l.pos] == '"' {
					inString = false
					l.pos++
					continue
				}
				l.pos++
				continue
			}

			if l.input[l.pos] == '"' {
				inString = true
				l.pos++
				continue
			}

			if (l.input[l.pos] == '{' || l.input[l.pos] == '[') && !inString {
				depth++
			}
			if (l.input[l.pos] == '}' || l.input[l.pos] == ']') && !inString {
				depth--
			}
			l.pos++
		}
	case 't':
		if l.pos+4 > len(l.input) {
			return nil, fmt.Errorf("scanRaw: unexpected end of input")
		}
		l.pos += 4
	case 'f':
		if l.pos+5 > len(l.input) {
			return nil, fmt.Errorf("scanRaw: unexpected end of input")
		}
		l.pos += 5
	case 'n':
		if l.pos+4 > len(l.input) {
			return nil, fmt.Errorf("scanRaw: unexpected end of input")
		}
		l.pos += 4
	default:
		if ch == '-' || (ch >= '0' && ch <= '9') {
			for l.pos < len(l.input) && ((l.input[l.pos] >= '0' && l.input[l.pos] <= '9') ||
				l.input[l.pos] == '.' || l.input[l.pos] == 'e' || l.input[l.pos] == 'E' ||
				l.input[l.pos] == '+' || l.input[l.pos] == '-') {
				l.pos++
			}
		}
	}

	return l.input[start:l.pos], nil
}

func bytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func (l *lexer) expectString() (string, error) {
	token, err := l.next()
	if err != nil {
		return "", err
	}
	if token.kind != tokString {
		return "", fmt.Errorf("expected 'string', got %v", token.kind)
	}

	return bytesToString(token.val), nil
}

func (l *lexer) expectNumber() ([]byte, error) {
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokNumber {
		return nil, fmt.Errorf("expected 'number', got %v", token.kind)
	}

	return token.val, nil
}
