package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func EncryptPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty!")
	}

	return GetMD5Hash(password), nil
}
