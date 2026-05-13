package openrtb

import "strconv"

func decodeBidRequest(l *lexer) (*BidRequest, error) {
	var br BidRequest
	if err := decodeBidRequestInto(l, &br); err != nil {
		return nil, err
	}
	return &br, nil
}

func decodeBidRequestInto(l *lexer, br *BidRequest) error {
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
			br.ID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "imp":
			br.Imp, err = decodeImpSlice(l, nil, nil)
			if err != nil {
				return err
			}
		case "at":
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
			br.AT = &n
		case "tmax":
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
			br.TMax = &n
		case "bcat":
			br.BCat, err = decodeStringSlice(l)
			if err != nil {
				return err
			}
		case "badv":
			br.BAdv, err = decodeStringSlice(l)
			if err != nil {
				return err
			}
		case "app":
			br.App = new(App)
			if err = decodeApp(l, br.App); err != nil {
				return err
			}
		case "site":
			br.Site = new(Site)
			if err = decodeSite(l, br.Site); err != nil {
				return err
			}
		case "user":
			br.User = new(User)
			if err = decodeUser(l, br.User); err != nil {
				return err
			}
		case "device":
			br.Device = new(Device)
			if err = decodeDevice(l, br.Device); err != nil {
				return err
			}
		case "ext":
			br.Ext, err = l.scanRaw()
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

func decodePublisher(l *lexer, pub *Publisher) error {
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
			pub.ID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "name":
			pub.Name, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "domain":
			pub.Domain, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "ext":
			pub.Ext, err = l.scanRaw()
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

func decodeSite(l *lexer, site *Site) error {
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
			site.ID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "name":
			site.Name, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "domain":
			site.Domain, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "page":
			site.Page, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "publisher":
			// If arena pre-allocated a Publisher, reuse it; otherwise allocate.
			if site.Publisher == nil {
				site.Publisher = new(Publisher)
			}
			if err = decodePublisher(l, site.Publisher); err != nil {
				return err
			}
		case "ext":
			site.Ext, err = l.scanRaw()
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

func decodeApp(l *lexer, app *App) error {
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
			app.ID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "name":
			app.Name, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "domain":
			app.Domain, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "bundle":
			app.Bundle, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "publisher":
			if app.Publisher == nil {
				app.Publisher = new(Publisher)
			}
			if err = decodePublisher(l, app.Publisher); err != nil {
				return err
			}
		case "ext":
			app.Ext, err = l.scanRaw()
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

func decodeDevice(l *lexer, device *Device) error {
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
		case "ua":
			device.UA, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "ip":
			device.IP, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "devicetype":
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
			device.DeviceType = &n
		case "make":
			device.Make, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "model":
			device.Model, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "os":
			device.OS, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "osv":
			device.OSV, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "language":
			device.Language, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "ext":
			device.Ext, err = l.scanRaw()
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

func decodeUser(l *lexer, user *User) error {
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
			user.ID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "buyeruid":
			user.BuyerUID, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "yob":
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
			user.Yob = &n
		case "gender":
			user.Gender, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "ext":
			user.Ext, err = l.scanRaw()
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

// decodeBanner fills banner from the lexer.
// formatBuf, if non-nil, is used as the backing array for banner.Format (arena path).
func decodeBanner(l *lexer, banner *Banner, formatBuf []Format) error {
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
		case "format":
			banner.Format, err = decodeFormatSlice(l, formatBuf)
			if err != nil {
				return err
			}
		case "w":
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
			banner.W = &n
		case "h":
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
			banner.H = &n
		case "mimes":
			banner.MIMEs, err = decodeStringSlice(l)
			if err != nil {
				return err
			}
		case "pos":
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
			banner.Pos = &n
		case "ext":
			banner.Ext, err = l.scanRaw()
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

// decodeImp fills imp from the lexer.
// formatBuf, if non-nil, is forwarded to decodeBanner (arena path).
func decodeImp(l *lexer, imp *Imp, formatBuf []Format) error {
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
			if imp.Banner == nil {
				imp.Banner = new(Banner)
			}
			if err = decodeBanner(l, imp.Banner, formatBuf); err != nil {
				return err
			}
		case "video":
			imp.Video = new(Video)
			if err = decodeVideo(l, imp.Video); err != nil {
				return err
			}
		case "native":
			imp.Native = new(Native)
			if err = decodeNative(l, imp.Native); err != nil {
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

func decodeNative(l *lexer, native *Native) error {
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
		case "request":
			native.Request, err = l.readStringVal()
			if err != nil {
				return err
			}
		case "ver":
			var s string
			s, err = l.readStringVal()
			if err != nil {
				return err
			}
			native.Ver = &s
		case "ext":
			native.Ext, err = l.scanRaw()
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

func decodeVideo(l *lexer, video *Video) error {
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
		case "mimes":
			video.MIMEs, err = decodeStringSlice(l)
			if err != nil {
				return err
			}
		case "minduration":
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
			video.MinDuration = &n
		case "maxduration":
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
			video.MaxDuration = &n
		case "protocols":
			video.Protocols, err = decodeIntSlice(l)
			if err != nil {
				return err
			}
		case "w":
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
			video.W = &n
		case "h":
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
			video.H = &n
		case "linearity":
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
			video.Linearity = &n
		case "skip":
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
			video.Skip = &n
		case "playbackmethod":
			video.PlaybackMethod, err = decodeIntSlice(l)
			if err != nil {
				return err
			}
		case "ext":
			video.Ext, err = l.scanRaw()
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

// decodeFormatSlice decodes a JSON array of Format objects.
// buf, if non-nil, is used as the initial backing array (arena path).
func decodeFormatSlice(l *lexer, buf []Format) ([]Format, error) {
	slice := buf
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

// decodeImpSlice decodes a JSON array of Imp objects.
// buf, if non-nil, is used as the initial backing array (arena path).
// formatBuf, if non-nil, is forwarded to each decodeImp call (arena path).
func decodeImpSlice(l *lexer, buf []Imp, formatBuf []Format) ([]Imp, error) {
	slice := buf
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
		if err := decodeImp(l, &slice[len(slice)-1], formatBuf); err != nil {
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
