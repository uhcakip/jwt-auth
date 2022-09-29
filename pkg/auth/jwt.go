package auth

import (
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

type JWT interface {
	GenerateAccessToken(subject string, expiresIn time.Duration) (accessToken string, err error)
	ParseAccessToken(accessToken string) (claims *jwt.RegisteredClaims, err error)
}

var jwtPrivateKey, jwtPublicKey []byte

type jwtAuth struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func SetupJWT() {
	var err error
	if jwtPrivateKey, err = ioutil.ReadFile(viper.GetString("filePath.jwtPrivateKey")); err != nil {
		log.Fatalln(err)
	}
	if jwtPublicKey, err = ioutil.ReadFile(viper.GetString("filePath.jwtPublicKey")); err != nil {
		log.Fatalln(err)
	}

	return
}

func NewJWT() (j JWT, err error) {
	ja := new(jwtAuth)

	if ja.privateKey, err = jwt.ParseECPrivateKeyFromPEM(jwtPrivateKey); err != nil {
		return
	}
	if ja.publicKey, err = jwt.ParseECPublicKeyFromPEM(jwtPublicKey); err != nil {
		return
	}

	return ja, nil
}

func (ja *jwtAuth) GenerateAccessToken(subject string, expiresIn time.Duration) (accessToken string, err error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   subject,
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(ja.privateKey)
	return
}

func (ja *jwtAuth) ParseAccessToken(accessToken string) (claims *jwt.RegisteredClaims, err error) {
	claims = new(jwt.RegisteredClaims)
	_, err = jwt.ParseWithClaims(accessToken, claims, func(*jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM(jwtPublicKey)
	})

	return
}
