package config

import (
	"github.com/spf13/viper"
	"log"
)

func Setup() {
	setDefaults()
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func setDefaults() {
	m := map[string]string{
		"app.env":                "local",
		"app.jwtKeyPath.private": "assets/jwt/private.key",
		"app.jwtKeyPath.public":  "assets/jwt/public.key",
	}

	for k, v := range m {
		viper.SetDefault(k, v)
	}
}
