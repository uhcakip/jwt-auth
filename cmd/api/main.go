package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"onboarding-jwt/pkg/auth"
	"onboarding-jwt/pkg/config"
	"time"
)

func main() {
	config.Setup()
	auth.Setup()

	var (
		err     error
		jwtAuth auth.JWT
	)

	if jwtAuth, err = auth.NewJWT(); err != nil {
		log.Println(err)
		return
	}

	var (
		accessToken string
		claims      *jwt.RegisteredClaims
	)

	if accessToken, err = jwtAuth.GenerateAccessToken("user_id_001", 2*time.Hour); err != nil {
		log.Println(err)
		return
	}
	if claims, err = jwtAuth.ParseAccessToken(accessToken); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(claims.Subject, claims.ExpiresAt)
}
