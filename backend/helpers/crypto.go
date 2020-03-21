package helpers

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// EncryptID create encryption based on string
func EncryptID(prefix string, val int) string {
	value := fmt.Sprintf("%s-%s", prefix, strconv.Itoa(val))
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// DecryptID to decrypt hash value
func DecryptID(prefix, val string) (int, error) {
	result, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return 0, err
	}

	r := strings.TrimPrefix(string(result), fmt.Sprintf("%s-", prefix))

	return strconv.Atoi(r)
}
