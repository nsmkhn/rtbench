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
			br.ID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "at":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = strconv.Atoi(string(val))
			if err != nil {
				return nil, err
			}
			br.AT = &n
		case "tmax":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = strconv.Atoi(string(val))
			if err != nil {
				return nil, err
			}
			br.TMax = &n
		case "bcat":
			br.BCat, err = decodeStringSlice(l)
			if err != nil {
				return nil, err
			}
		case "badv":
			br.BAdv, err = decodeStringSlice(l)
			if err != nil {
				return nil, err
			}
		case "app":
			br.App, err = decodeApp(l)
			if err != nil {
				return nil, err
			}
		case "site":
			br.Site, err = decodeSite(l)
			if err != nil {
				return nil, err
			}
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
			break
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return &br, nil
}

func decodePublisher(l *lexer) (*Publisher, error) {
	var pub Publisher

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
			return nil, fmt.Errorf("expected 'string', got %v", token.kind)
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
			pub.ID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "name":
			pub.Name, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "domain":
			pub.Domain, err = l.expectString()
			if err != nil {
				return nil, err
			}
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
			break
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return &pub, nil
}

func decodeSite(l *lexer) (*Site, error) {
	var site Site

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
			return nil, fmt.Errorf("expected 'string', got %v", token.kind)
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
			site.ID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "name":
			site.Name, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "domain":
			site.Domain, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "page":
			site.Page, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "publisher":
			site.Publisher, err = decodePublisher(l)
			if err != nil {
				return nil, err
			}
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
			break
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return &site, nil
}

func decodeApp(l *lexer) (*App, error) {
	var app App

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
			return nil, fmt.Errorf("expected 'string', got %v", token.kind)
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
			app.ID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "name":
			app.Name, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "domain":
			app.Domain, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "bundle":
			app.Bundle, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "publisher":
			app.Publisher, err = decodePublisher(l)
			if err != nil {
				return nil, err
			}
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
			break
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return &app, nil
}

func decodeStringSlice(l *lexer) ([]string, error) {
	var slice []string
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokLBracket {
		return nil, fmt.Errorf("expected '[', got %v", token.kind)
	}

	for {
		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			return slice, nil
		}
		if token.kind != tokString {
			return nil, fmt.Errorf("expected 'string', got %v", token.kind)
		}
		slice = append(slice, string(token.val))

		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			return slice, nil
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}
}

func decodeIntSlice(l *lexer) ([]int, error) {
	var slice []int
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokLBracket {
		return nil, fmt.Errorf("expected '[', got %v", token.kind)
	}

	for {
		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			return slice, nil
		}
		if token.kind != tokNumber {
			return nil, fmt.Errorf("expected 'number', got %v", token.kind)
		}
		var n int
		n, err = strconv.Atoi(string(token.val))
		if err != nil {
			return nil, err
		}
		slice = append(slice, n)

		token, err = l.next()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			return slice, nil
		}
		if token.kind != tokComma {
			return nil, fmt.Errorf("expected ',', got %v", token.kind)
		}
	}
}
