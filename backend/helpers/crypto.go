package helpers

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// PrefixCursor is a constant to prefix encryptor
const PrefixCursor = "cursor-"

// EncryptID create encryption based on string
func EncryptID(val string) string {
	value := fmt.Sprintf("%s-%s", PrefixCursor, val)
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// DecryptID to decrypt hash value
func DecryptID(val string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(string(result), PrefixCursor), nil
}
