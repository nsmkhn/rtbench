package openrtb

import (
	"sync"
	"unsafe"
)

// Arena holds a BidRequest and the most commonly heap-allocated sub-objects in
// a single allocation. BidRequest MUST be the first field: ReleaseArena casts
// *BidRequest back to *Arena via unsafe.Pointer, which is only valid when the
// two share the same base address.
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

// ParseFastArena is like ParseFast but obtains a pooled Arena to reduce heap
// allocations. Call ReleaseArena(br) when the BidRequest is no longer needed.
//
// String fields in the returned BidRequest point into data — do not modify
// data while the BidRequest is in use.
func ParseFastArena(data []byte) (*BidRequest, error) {
	arena := arenaPool.Get().(*Arena)

	// Zero only the fields that will be written into; impBuf is zeroed on use
	// because callers pass arena.impBuf[:0] (a zero-length slice).
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

// ReleaseArena returns br and its owning Arena to the pool.
// br must have been obtained from ParseFastArena; behaviour is undefined otherwise.
func ReleaseArena(br *BidRequest) {
	// Safe because BidRequest is the first field of Arena (offset 0).
	arenaPool.Put((*Arena)(unsafe.Pointer(br)))
}

// decodeBidRequestArena is a copy of decodeBidRequestInto that uses pre-allocated
// sub-objects from the arena rather than allocating with new().
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
			// Use arena.impBuf as the backing array so the []Imp slice header is
			// the only thing that escapes to the heap (if len > 8, append allocates
			// a new backing array, which is unavoidable).
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
