package auth

import (
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"onboarding-jwt/pkg/helper"
	"time"
)

type JWT interface {
	GenerateAccessToken(subject string, expiry time.Duration) (accessToken string, err error)
	ParseAccessToken(accessToken string) (claims *jwt.RegisteredClaims, err error)
}

var jwtPrivateKey, jwtPublicKey []byte

type jwtAuth struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func Setup() {
	var (
		err      error
		rootPath string
	)

	if rootPath, err = helper.GetRootDirPath(); err != nil {
		log.Fatalln(err)
	}

	privateKeyPath := rootPath + "/" + viper.GetString("app.jwtKeyPath.private")
	publicKeyPath := rootPath + "/" + viper.GetString("app.jwtKeyPath.public")

	if jwtPrivateKey, err = ioutil.ReadFile(privateKeyPath); err != nil {
		log.Fatalln(err)
	}
	if jwtPublicKey, err = ioutil.ReadFile(publicKeyPath); err != nil {
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

func (ja *jwtAuth) GenerateAccessToken(subject string, expiry time.Duration) (accessToken string, err error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
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
