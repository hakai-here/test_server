package utils

import (
	"demoproject/api/constant"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

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
