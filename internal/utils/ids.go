package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateSessionToken(username string) string {
	shaAlgorithm := sha256.New()

	shaAlgorithm.Write([]byte(username))

	hashedUsername := hex.EncodeToString(shaAlgorithm.Sum(nil))
	tokenId := uuid.New().String()

	return hashedUsername[:12] + "." + tokenId
}
