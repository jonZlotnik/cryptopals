package main

import (
	"encoding/hex"
	"log"
)

// FixedXOR ...
func FixedXOR(a, b string) string {

	aBytes, errA := hex.DecodeString(a)
	bBytes, errB := hex.DecodeString(b)

	if errA != nil || errB != nil {
		log.Fatal(errA, errB)
	}
	if len(aBytes) != len(bBytes) {
		log.Fatal()
	}

	outBytes := make([]byte, len(aBytes))

	for i := 0; i < len(outBytes); i++ {
		outBytes[i] = aBytes[i] ^ bBytes[i]
	}

	return hex.EncodeToString(outBytes)
}
