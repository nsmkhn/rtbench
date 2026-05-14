package openrtb

import (
	"sync"
	"unsafe"
)

type Arena struct {
	BidRequest // MUST remain first
	site       Site
	app        App
	device     Device
	user       User
	impBuf     [8]Imp
}

var arenaPool = sync.Pool{
	New: func() any { return new(Arena) },
}

func ParseFastArena(data []byte) (*BidRequest, error) {
	arena := arenaPool.Get().(*Arena)

	arena.BidRequest = BidRequest{}
	arena.site = Site{}
	arena.app = App{}
	arena.device = Device{}
	arena.user = User{}

	l := newLexer(data)
	if err := decodeBidRequestArena(l, arena); err != nil {
		arenaPool.Put(arena)
		return nil, err
	}
	return &arena.BidRequest, nil
}

func ReleaseArena(br *BidRequest) {
	arenaPool.Put((*Arena)(unsafe.Pointer(br)))
}

func decodeBidRequestArena(l *lexer, arena *Arena) error {
	br := &arena.BidRequest
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
			br.Imp, err = decodeImpSlice(l, arena.impBuf[:0], nil)
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
		case "site":
			br.Site = &arena.site
			if err = decodeSite(l, br.Site); err != nil {
				return err
			}
		case "app":
			br.App = &arena.app
			if err = decodeApp(l, br.App); err != nil {
				return err
			}
		case "device":
			br.Device = &arena.device
			if err = decodeDevice(l, br.Device); err != nil {
				return err
			}
		case "user":
			br.User = &arena.user
			if err = decodeUser(l, br.User); err != nil {
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
