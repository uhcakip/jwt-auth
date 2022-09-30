package web

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"jwt-auth/pkg/oauth"
	"net/http"
)

func OAuthLogin(provider oauth.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		oauthConfig := provider.GetConfig() // TODO: check zero value
		state := uuid.New().String()        // TODO: store state value to db or redis
		c.Redirect(http.StatusSeeOther, oauthConfig.AuthCodeURL(state))
	}
}
