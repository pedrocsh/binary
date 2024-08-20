package main

import (
	"encoding/hex"
)

func DecodeHex(input string) string {
	src := []byte(input)
	dst := make([]byte, len(src))
	decoded, _ := hex.Decode(dst, src)

	return string(dst[:decoded])
}
