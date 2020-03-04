package helpers

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// PrefixCursor is a constant to prefix encryptor
const PrefixCursor = "cursor-"

// EncryptCursor create encryption based on string
func EncryptCursor(val string) string {
	value := fmt.Sprintf("%s-%s", PrefixCursor, val)
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// DecryptCursor to decrypt hash value
func DecryptCursor(val string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(string(result), PrefixCursor), nil
}
