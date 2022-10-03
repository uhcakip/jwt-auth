package main

import (
	"jwt-auth/pkg/auth"
	"jwt-auth/pkg/config"
	"jwt-auth/router"
)

func main() {
	/*
		var oauth oauth.Provider
		c := oauth.GetConfig()
		if reflect.ValueOf(c).IsZero() {
			println("zero")
		}
		if reflect.ValueOf(c).IsNil() {
			println("nil")
		}
		return
	*/

	config.Setup()
	auth.SetupJWT()
	router.Route()
}
