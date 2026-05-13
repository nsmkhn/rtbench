//go:build !arm64

package openrtb

func indexStopByte(b []byte) int {
	for i, c := range b {
		if c == '"' || c == '\\' {
			return i
		}
	}
	return len(b)
}
