package utils

import (
	"crypto/md5"
	"fmt"
)

func GetHash(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func IsHashValid(hash, key string) bool {
	if hash == fmt.Sprintf("%x", md5.Sum([]byte(key))) {
		return true
	} else {
		return false
	}
}
