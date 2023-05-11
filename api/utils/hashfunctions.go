package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"
)

var salt = []byte(viper.GetString("SALTVALUE"))

func HashArgon2(password string) string {
	time := uint32(2)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLength := uint32(32)
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	return fmt.Sprintf("%x", hash)
}
