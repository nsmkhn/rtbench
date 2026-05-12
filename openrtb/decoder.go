package openrtb

import (
	"strconv"
)

func decodeBidRequest(l *lexer) (*BidRequest, error) {
	var br BidRequest

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "id":
			br.ID, err = l.readStringVal()
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
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			br.AT = &n
		case "tmax":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
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
			br.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &br, nil
}

func decodePublisher(l *lexer) (*Publisher, error) {
	var pub Publisher

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "id":
			pub.ID, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "name":
			pub.Name, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "domain":
			pub.Domain, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "ext":
			pub.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &pub, nil
}

func decodeSite(l *lexer) (*Site, error) {
	var site Site

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "id":
			site.ID, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "name":
			site.Name, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "domain":
			site.Domain, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "page":
			site.Page, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "publisher":
			site.Publisher, err = decodePublisher(l)
			if err != nil {
				return nil, err
			}
		case "ext":
			site.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &site, nil
}

func decodeApp(l *lexer) (*App, error) {
	var app App

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "id":
			app.ID, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "name":
			app.Name, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "domain":
			app.Domain, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "bundle":
			app.Bundle, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "publisher":
			app.Publisher, err = decodePublisher(l)
			if err != nil {
				return nil, err
			}
		case "ext":
			app.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &app, nil
}

func decodeDevice(l *lexer) (*Device, error) {
	var device Device

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "ua":
			device.UA, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "ip":
			device.IP, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "devicetype":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			device.DeviceType = &n
		case "make":
			device.Make, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "model":
			device.Model, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "os":
			device.OS, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "osv":
			device.OSV, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "language":
			device.Language, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "ext":
			device.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &device, nil
}

func decodeUser(l *lexer) (*User, error) {
	var user User

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "id":
			user.ID, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "buyeruid":
			user.BuyerUID, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "yob":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			user.Yob = &n
		case "gender":
			user.Gender, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "ext":
			user.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &user, nil
}

func decodeFormat(l *lexer, format *Format) error {
	if err := l.skipOpen(); err != nil {
		return err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return err
		}
		if err = l.skipColon(); err != nil {
			return err
		}

		switch string(key) {
		case "w":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return err
			}
			format.W, err = parseIntBytes(val)
			if err != nil {
				return err
			}
		case "h":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return err
			}
			format.H, err = parseIntBytes(val)
			if err != nil {
				return err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return err
		}
		if done {
			break
		}
	}

	return nil
}

func decodeBanner(l *lexer) (*Banner, error) {
	var banner Banner

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "format":
			banner.Format, err = decodeFormatSlice(l)
			if err != nil {
				return nil, err
			}
		case "w":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			banner.W = &n
		case "h":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
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
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			banner.Pos = &n
		case "ext":
			banner.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &banner, nil
}

func decodeImp(l *lexer, imp *Imp) error {
	if err := l.skipOpen(); err != nil {
		return err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return err
		}
		if err = l.skipColon(); err != nil {
			return err
		}

		switch string(key) {
		case "id":
			imp.ID, err = l.readStringVal()
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
			val, err = l.readNumberBytes()
			if err != nil {
				return err
			}
			var f float64
			f, err = strconv.ParseFloat(bytesToString(val), 64)
			if err != nil {
				return err
			}
			imp.BidFloor = &f
		case "bidfloorcur":
			imp.BidFloorCur, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "secure":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return err
			}
			imp.Secure = &n
		case "ext":
			imp.Ext, err = l.scanRaw()
			if err != nil {
				return err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return err
		}
		if done {
			break
		}
	}

	return nil
}

func decodeNative(l *lexer) (*Native, error) {
	var native Native

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "request":
			native.Request, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
		case "ver":
			var s string
			s, err = l.readStringVal()
			if err != nil {
				return nil, err
			}
			native.Ver = &s
		case "ext":
			native.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &native, nil
}

func decodeVideo(l *lexer) (*Video, error) {
	var video Video

	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == '}' {
			l.pos++
			break
		}

		key, err := l.readKey()
		if err != nil {
			return nil, err
		}
		if err = l.skipColon(); err != nil {
			return nil, err
		}

		switch string(key) {
		case "mimes":
			video.MIMEs, err = decodeStringSlice(l)
			if err != nil {
				return nil, err
			}
		case "minduration":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			video.MinDuration = &n
		case "maxduration":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
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
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			video.W = &n
		case "h":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			video.H = &n
		case "linearity":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
			if err != nil {
				return nil, err
			}
			video.Linearity = &n
		case "skip":
			var val []byte
			val, err = l.readNumberBytes()
			if err != nil {
				return nil, err
			}
			var n int
			n, err = parseIntBytes(val)
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
			video.Ext, err = l.scanRaw()
			if err != nil {
				return nil, err
			}
		default:
			if _, err = l.scanRaw(); err != nil {
				return nil, err
			}
		}

		var done bool
		done, err = l.readSep('}')
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	return &video, nil
}

func decodeStringSlice(l *lexer) ([]string, error) {
	var slice []string
	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == ']' {
			l.pos++
			return slice, nil
		}

		s, err := l.readStringVal()
		if err != nil {
			return nil, err
		}
		slice = append(slice, s)

		done, err := l.readSep(']')
		if err != nil {
			return nil, err
		}
		if done {
			return slice, nil
		}
	}
}

func decodeIntSlice(l *lexer) ([]int, error) {
	var slice []int
	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == ']' {
			l.pos++
			return slice, nil
		}

		val, err := l.readNumberBytes()
		if err != nil {
			return nil, err
		}
		n, err := parseIntBytes(val)
		if err != nil {
			return nil, err
		}
		slice = append(slice, n)

		done, err := l.readSep(']')
		if err != nil {
			return nil, err
		}
		if done {
			return slice, nil
		}
	}
}

func decodeFormatSlice(l *lexer) ([]Format, error) {
	var slice []Format
	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == ']' {
			l.pos++
			return slice, nil
		}

		slice = append(slice, Format{})
		if err := decodeFormat(l, &slice[len(slice)-1]); err != nil {
			return nil, err
		}

		done, err := l.readSep(']')
		if err != nil {
			return nil, err
		}
		if done {
			return slice, nil
		}
	}
}

func decodeImpSlice(l *lexer) ([]Imp, error) {
	var slice []Imp
	if err := l.skipOpen(); err != nil {
		return nil, err
	}

	for {
		l.skipWS()
		if l.pos < len(l.input) && l.input[l.pos] == ']' {
			l.pos++
			return slice, nil
		}

		slice = append(slice, Imp{})
		if err := decodeImp(l, &slice[len(slice)-1]); err != nil {
			return nil, err
		}

		done, err := l.readSep(']')
		if err != nil {
			return nil, err
		}
		if done {
			return slice, nil
		}
	}
}
