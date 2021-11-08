package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/hiagomf/bank-api/server/config"
)

// EncryptPassword - Encripta o dado
func EncryptPassword(password string) (hash string, err error) {
	secrect := config.GetConfig().Secrets[0]

	h := hmac.New(sha256.New, []byte(secrect))
	_, err = h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hash = hex.EncodeToString(h.Sum(nil))
	return
}

// CheckEncryptedPassword - verifica se os dois dados são iguais
func CheckEncryptedPassword(passwordIn string, passwordBD string) (ok bool, err error) {
	secrect := config.GetConfig().Secrets[0]

	h := hmac.New(sha256.New, []byte(secrect))
	_, err = h.Write([]byte(passwordIn))
	if err != nil {
		return false, err
	}

	if hex.EncodeToString(h.Sum(nil)) == passwordBD {
		return true, err
	}
	return
}
