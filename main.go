package main

import "fmt"

func main() {
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hex_decoded := DecodeHex(str)
	base64_encoded := EncodeBase64(hex_decoded)
	fmt.Println("String: " + hex_decoded)
	fmt.Println("Base64: " + base64_encoded + "=")
}
