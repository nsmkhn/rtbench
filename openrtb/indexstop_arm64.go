//go:build arm64

package openrtb

// indexStopByte returns the index of the first '"' (0x22) or '\' (0x5C) byte
// in b, or len(b) if neither is found.
//
//go:noescape
func indexStopByte(b []byte) int
