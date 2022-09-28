package config

import (
	"github.com/spf13/viper"
	"log"
)

func Setup() {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
