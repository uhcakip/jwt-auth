package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"jwt-auth/pkg/oauth"
	"jwt-auth/router/api"
	"jwt-auth/router/web"
	"log"
	"net/http"
)

var facebook oauth.Provider

func Route() {
	facebook = oauth.NewProvider(oauth.ProviderNameFacebook)

	r := gin.Default()
	r.Static("./assets/css", "./assets/css")
	r.LoadHTMLGlob("./views/*")

	routeWeb(r)
	routeAPI(r)

	if err := r.RunTLS("", viper.GetString("filePath.sslCert"), viper.GetString("filePath.sslKey")); err != nil {
		log.Fatalln(err)
	}
}

func routeWeb(r *gin.Engine) {
	/*
		GET /
		GET /oauth/facebook/login
		GET /oauth/facebook/callback
	*/

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.gohtml", nil)
	})

	oauthFacebookRoute := r.Group("oauth/facebook")
	{
		oauthFacebookRoute.GET("login", web.OAuthLogin(facebook))
		oauthFacebookRoute.GET("callback", web.OAuthCallback(facebook))
	}
}

func routeAPI(r *gin.Engine) {
	/*
			POST /oaut


		m
	*/

	r.POST("oauth/login", api.OAuthLogin(facebook))
}
