package web

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/pkg/oauth"
	"net/http"
)

type OAuthCallbackPayload struct {
	Code  string `form:"code"  binding:"required"`
	State string `form:"state" binding:"required"`
}

func OAuthCallback(provider oauth.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload OAuthCallbackPayload
		if err := c.ShouldBindQuery(&payload); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// state, code can be used only for once
		c.HTML(http.StatusOK, "views/callback.gohtml", nil)
		return
	}
}
