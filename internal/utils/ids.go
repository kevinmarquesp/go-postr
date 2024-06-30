package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

// TODO: It should not depend on the `username` field to generate a new session token id.

func GenerateSessionToken(username string) string {
	shaAlgorithm := sha256.New()

	shaAlgorithm.Write([]byte(username))

	hashedUsername := hex.EncodeToString(shaAlgorithm.Sum(nil))
	tokenId := uuid.New().String()

	return hashedUsername[:12] + "." + tokenId
}
