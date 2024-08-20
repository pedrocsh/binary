package main

import "encoding/base64"

func EncodeBase64(input string) string {
  src := []byte(input)
  encoded := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
  base64.StdEncoding.Encode(encoded, src)

  return string(encoded)
}
