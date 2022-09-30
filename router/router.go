package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"jwt-auth/pkg/oauth"
	"jwt-auth/router/web"
	"log"
	"net/http"
)

func Route() {
	r := gin.Default()
	r.Static("./assets/css", "./assets/css")
	r.LoadHTMLGlob("./views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.gohtml", nil)
	})

	routeWeb(r)

	if err := r.RunTLS("", viper.GetString("filePath.sslCert"), viper.GetString("filePath.sslKey")); err != nil {
		log.Fatalln(err)
	}
}

func routeWeb(r *gin.Engine) {
	/*
		/oauth/facebook/login
		/oauth/facebook/callback
	*/

	providerFB := oauth.NewProvider(oauth.ProviderNameFacebook)
	oauthFb := r.Group("oauth/facebook")
	{
		oauthFb.GET("login", web.OAuthLogin(providerFB))
		oauthFb.GET("callback", web.OAuthCallback(providerFB))
	}
}
