package openrtb

import (
	"fmt"
	"unsafe"
)

type lexer struct {
	input []byte
	pos   int
}

func newLexer(data []byte) *lexer {
	return &lexer{input: data}
}

func (l *lexer) skipWS() {
	for l.pos < len(l.input) &&
		(l.input[l.pos] == ' ' || l.input[l.pos] == '\n' || l.input[l.pos] == '\t' || l.input[l.pos] == '\r') {
		l.pos++
	}
}

func (l *lexer) readKey() ([]byte, error) {
	l.skipWS()

	if l.pos >= len(l.input) {
		return nil, fmt.Errorf("readKey: unexpected end of input")
	}

	if l.input[l.pos] != '"' {
		return nil, fmt.Errorf("readKey: expected '\"', got %q", l.input[l.pos])
	}

	l.pos++
	start := l.pos

	for l.pos < len(l.input) && l.input[l.pos] != '"' {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return nil, fmt.Errorf("readKey: unterminated key")
	}

	key := l.input[start:l.pos]
	l.pos++
	return key, nil
}

func (l *lexer) skipColon() error {
	l.skipWS()
	if l.pos >= len(l.input) || l.input[l.pos] != ':' {
		return fmt.Errorf("skipColon: expected ':'")
	}

	l.pos++

	return nil
}

func (l *lexer) skipOpen() error {
	l.skipWS()
	if l.pos >= len(l.input) || (l.input[l.pos] != '[' && l.input[l.pos] != '{') {
		return fmt.Errorf("skipOpen: expected '{' or '['")
	}

	l.pos++

	return nil
}

func (l *lexer) readSep(close byte) (bool, error) {
	l.skipWS()

	if l.pos >= len(l.input) {
		return false, fmt.Errorf("readSep: unexpected end of input")
	}

	if l.input[l.pos] == close {
		l.pos++
		return true, nil
	}
	if l.input[l.pos] == ',' {
		l.pos++
		return false, nil
	}

	return false, fmt.Errorf("readSep: expected ',' or %q, got %q", close, l.input[l.pos])
}

func (l *lexer) scanString() ([]byte, error) {
	l.pos++ // skip opening '"'
	start := l.pos
	var buf []byte

	for {
		remaining := l.input[l.pos:]
		n := indexStopByte(remaining)

		if n == len(remaining) {
			return nil, fmt.Errorf("lexer: unterminated string")
		}

		ch := l.input[l.pos+n]

		if ch == '"' {
			if buf == nil {
				val := l.input[start : l.pos+n]
				l.pos += n + 1
				return val, nil
			}
			buf = append(buf, remaining[:n]...)
			l.pos += n + 1
			return buf, nil
		}

		// ch == '\\'
		if buf == nil {
			buf = append(buf, l.input[start:l.pos+n]...)
		} else {
			buf = append(buf, remaining[:n]...)
		}
		l.pos += n + 1

		if l.pos >= len(l.input) {
			return nil, fmt.Errorf("lexer: unterminated escape")
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
			return nil, fmt.Errorf("lexer: unsupported escape sequence %q", l.input[l.pos])
		}
		l.pos++
	}
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

			if l.input[l.pos] == '{' || l.input[l.pos] == '[' {
				depth++
			}
			if l.input[l.pos] == '}' || l.input[l.pos] == ']' {
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

func (l *lexer) readStringVal() (string, error) {
	l.skipWS()
	if l.pos >= len(l.input) || l.input[l.pos] != '"' {
		return "", fmt.Errorf("readStringVal: expected '\"'")
	}
	val, err := l.scanString()
	if err != nil {
		return "", err
	}
	return bytesToString(val), nil
}

func (l *lexer) readNumberBytes() ([]byte, error) {
	l.skipWS()
	if l.pos >= len(l.input) {
		return nil, fmt.Errorf("readNumberBytes: unexpected end of input")
	}
	ch := l.input[l.pos]
	if ch != '-' && (ch < '0' || ch > '9') {
		return nil, fmt.Errorf("readNumberBytes: expected number, got %q", ch)
	}
	start := l.pos
	for l.pos < len(l.input) && ((l.input[l.pos] >= '0' && l.input[l.pos] <= '9') ||
		l.input[l.pos] == '.' || l.input[l.pos] == 'e' || l.input[l.pos] == 'E' ||
		l.input[l.pos] == '+' || l.input[l.pos] == '-') {
		l.pos++
	}
	return l.input[start:l.pos], nil
}

func parseIntBytes(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, fmt.Errorf("parseIntBytes: empty input")
	}
	neg := b[0] == '-'
	if neg {
		b = b[1:]
	}
	n := 0
	for _, c := range b {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("parseIntBytes: invalid byte %q", c)
		}
		n = n*10 + int(c-'0')
	}
	if neg {
		return -n, nil
	}
	return n, nil
}
