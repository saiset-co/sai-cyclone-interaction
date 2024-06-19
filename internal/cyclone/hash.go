package cyclone

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-interaction/logger"
)

func CreateHash(text string) (string, error) {
	return createSHA256Hash(text)
}

func createSHA256Hash(text string) (string, error) {
	hasher := sha256.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		logger.Logger.Error("createSHA256Hash", zap.Error(err))
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
