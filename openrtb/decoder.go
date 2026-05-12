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
		case "imp":
			br.Imp, err = decodeImpSlice(l)
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
		case "user":
			br.User, err = decodeUser(l)
			if err != nil {
				return nil, err
			}
		case "device":
			br.Device, err = decodeDevice(l)
			if err != nil {
				return nil, err
			}
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			br.Ext = raw
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
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			pub.Ext = raw
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
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			site.Ext = raw
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
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			app.Ext = raw
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

func decodeDevice(l *lexer) (*Device, error) {
	var device Device

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
		case "ua":
			device.UA, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "ip":
			device.IP, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "devicetype":
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
			device.DeviceType = &n

		case "make":
			device.Make, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "model":
			device.Model, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "os":
			device.OS, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "osv":
			device.OSV, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "language":
			device.Language, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			device.Ext = raw
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

	return &device, nil
}

func decodeUser(l *lexer) (*User, error) {
	var user User

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
			user.ID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "buyeruid":
			user.BuyerUID, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "yob":
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
			user.Yob = &n
		case "gender":
			user.Gender, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			user.Ext = raw
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

	return &user, nil
}

func decodeFormat(l *lexer, format *Format) error {
	token, err := l.next()
	if err != nil {
		return err
	}
	if token.kind != tokLBrace {
		return fmt.Errorf("expected '{', got %v", token.kind)
	}

	for {
		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind == tokRBrace {
			break
		}

		key := token
		if key.kind != tokString {
			return fmt.Errorf("expected 'string', got %v", token.kind)
		}

		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind != tokColon {
			return fmt.Errorf("expected ':', got %v", token.kind)
		}

		switch string(key.val) {
		case "w":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return err
			}

			format.W, err = strconv.Atoi(string(val))
			if err != nil {
				return err
			}
		case "h":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return err
			}

			format.H, err = strconv.Atoi(string(val))
			if err != nil {
				return err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return err
			}
		}

		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind == tokRBrace {
			break
		}
		if token.kind != tokComma {
			return fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return nil
}

func decodeBanner(l *lexer) (*Banner, error) {
	var banner Banner

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
		case "format":
			banner.Format, err = decodeFormatSlice(l)
			if err != nil {
				return nil, err
			}

		case "w":
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
			banner.W = &n
		case "h":
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
			banner.H = &n
		case "mimes":
			banner.MIMEs, err = decodeStringSlice(l)
			if err != nil {
				return nil, err
			}
		case "pos":
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
			banner.Pos = &n
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			banner.Ext = raw
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

	return &banner, nil
}

func decodeImp(l *lexer, imp *Imp) error {
	token, err := l.next()
	if err != nil {
		return err
	}
	if token.kind != tokLBrace {
		return fmt.Errorf("expected '{', got %v", token.kind)
	}

	for {
		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind == tokRBrace {
			break
		}

		key := token
		if key.kind != tokString {
			return fmt.Errorf("expected 'string', got %v", token.kind)
		}

		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind != tokColon {
			return fmt.Errorf("expected ':', got %v", token.kind)
		}

		switch string(key.val) {
		case "id":
			imp.ID, err = l.expectString()
			if err != nil {
				return err
			}

		case "banner":
			imp.Banner, err = decodeBanner(l)
			if err != nil {
				return err
			}
		case "video":
			imp.Video, err = decodeVideo(l)
			if err != nil {
				return err
			}
		case "native":
			imp.Native, err = decodeNative(l)
			if err != nil {
				return err
			}
		case "bidfloor":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return err
			}
			var f float64
			f, err = strconv.ParseFloat(string(val), 64)
			if err != nil {
				return err
			}
			imp.BidFloor = &f
		case "bidfloorcur":
			imp.BidFloorCur, err = l.expectString()
			if err != nil {
				return err
			}
		case "secure":
			var val []byte
			val, err = l.expectNumber()
			if err != nil {
				return err
			}
			var n int
			n, err = strconv.Atoi(string(val))
			if err != nil {
				return err
			}
			imp.Secure = &n
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return err
			}
			imp.Ext = raw
		default:
			if _, err = l.scanRaw(); err != nil {
				return err
			}
		}

		token, err = l.next()
		if err != nil {
			return err
		}
		if token.kind == tokRBrace {
			break
		}
		if token.kind != tokComma {
			return fmt.Errorf("expected ',', got %v", token.kind)
		}
	}

	return nil
}

func decodeNative(l *lexer) (*Native, error) {
	var native Native

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
		case "request":
			native.Request, err = l.expectString()
			if err != nil {
				return nil, err
			}
		case "ver":
			var s string
			s, err = l.expectString()
			if err != nil {
				return nil, err
			}
			native.Ver = &s
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			native.Ext = raw
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

	return &native, nil
}

func decodeVideo(l *lexer) (*Video, error) {
	var video Video

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
		case "mimes":
			video.MIMEs, err = decodeStringSlice(l)
			if err != nil {
				return nil, err
			}
		case "minduration":
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
			video.MinDuration = &n
		case "maxduration":
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
			video.MaxDuration = &n

		case "protocols":
			video.Protocols, err = decodeIntSlice(l)
			if err != nil {
				return nil, err
			}
		case "w":
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
			video.W = &n
		case "h":
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
			video.H = &n
		case "linearity":
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
			video.Linearity = &n
		case "skip":
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
			video.Skip = &n
		case "playbackmethod":
			video.PlaybackMethod, err = decodeIntSlice(l)
			if err != nil {
				return nil, err
			}
		case "ext":
			var raw []byte
			raw, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
			video.Ext = raw
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

	return &video, nil
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

func decodeFormatSlice(l *lexer) ([]Format, error) {
	var slice []Format
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokLBracket {
		return nil, fmt.Errorf("expected '[', got %v", token.kind)
	}

	for {
		token, err = l.peek()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			l.next()
			return slice, nil
		}

		slice = append(slice, Format{})
		if err = decodeFormat(l, &slice[len(slice)-1]); err != nil {
			return nil, err
		}

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

func decodeImpSlice(l *lexer) ([]Imp, error) {
	var slice []Imp
	token, err := l.next()
	if err != nil {
		return nil, err
	}
	if token.kind != tokLBracket {
		return nil, fmt.Errorf("expected '[', got %v", token.kind)
	}

	for {
		token, err = l.peek()
		if err != nil {
			return nil, err
		}
		if token.kind == tokRBracket {
			l.next()
			return slice, nil
		}

		slice = append(slice, Imp{})
		if err = decodeImp(l, &slice[len(slice)-1]); err != nil {
			return nil, err
		}

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
