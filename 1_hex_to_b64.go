package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// HexToB64 ...
func HexToB64(hexInput string) string {
	byteArr, err := hex.DecodeString(hexInput)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(byteArr)
}
