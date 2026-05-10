package openrtb

import (
	"fmt"
	"strconv"
)

func decodeBidRequest(l *lexer) (*BidRequest, error) {
	var br BidRequest

	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokLBrace {
		return nil, fmt.Errorf("expected '{', got %v", token.kind)
	}

	for {
		token, err = l.next()
		if err != nil {
			return nil, err
		}

		if token.kind == tokRBrace {
			break
		}
		key := token
		if key.kind != tokString {
			return nil, fmt.Errorf("expected 'string' key, got %v", key.kind)
		}

		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind != tokColon {
			return nil, fmt.Errorf("expected ':', got %v", token.kind)
		}

		switch string(key.val) {
		case "id":
			token, err = l.next()
			if err != nil {
				return nil, err
			}
			if token.kind != tokString {
				return nil, fmt.Errorf("expected 'string', got %v", token.kind)
			}
			br.ID = string(token.val)
		case "at":
			token, err = l.next()
			if err != nil {
				return nil, err
			}
			if token.kind != tokNumber {
				return nil, fmt.Errorf("expected 'number', got %v", token.kind)
			}
			var n int
			n, err = strconv.Atoi(string(token.val))
			if err != nil {
				return nil, err
			}
			br.AT = &n
		case "tmax":
			token, err = l.next()
			if err != nil {
				return nil, err
			}
			if token.kind != tokNumber {
				return nil, fmt.Errorf("expected 'number', got %v", token.kind)
			}
			var n int
			n, err = strconv.Atoi(string(token.val))
			if err != nil {
				return nil, err
			}
			br.TMax = &n
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBrace {
			return &br, nil
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return &br, nil
}
