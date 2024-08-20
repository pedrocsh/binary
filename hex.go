package main

import (
	"encoding/hex"
	"fmt"
)

func DecodeHex(input string) string {
	src := []byte(input)
	dst := make([]byte, len(src))
	decoded, _ := hex.Decode(dst, src)
	fmt.Println(dst)

	return string(dst[:decoded])
}
