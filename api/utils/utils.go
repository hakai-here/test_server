package utils

import (
	"crypto/rand"
	"demoproject/api/constant"
	"fmt"
	"log"
	"unsafe"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ImportEnv() { // setting up the envionment variables
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 3000)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok { // Config file not found Ignoring
		} else {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	for _, element := range constant.ENV {
		if viper.GetString(element) == "" {
			log.Panicln(fmt.Errorf("envionment Variables are not present : %s", element))
		}
	}

}

func GenerateUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic("Failed to generate UUID")
	}
	return uuid.String()
}

func RandStr() string {
	size := 50 // specifing the length of the string
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
