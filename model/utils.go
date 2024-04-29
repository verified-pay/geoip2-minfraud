package model

import (
	"crypto/sha256"
	"encoding/hex"
)

func getSha256(input []byte) string {
	h := sha256.New()
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil))
}
