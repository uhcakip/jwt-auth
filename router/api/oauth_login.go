package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"jwt-auth/pkg/auth"
	"jwt-auth/pkg/helper"
	"jwt-auth/pkg/oauth"
	"jwt-auth/pkg/response"
	"log"
	"net/http"
	"time"
)

type OAuthLoginRequest struct {
	Code  string `json:"code"  binding:"required"`
	State string `json:"state" binding:"required"`
}

type OAuthLoginResponse struct {
	AccessToken string `json:"access_token"`
}

func OAuthLogin(provider oauth.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		apiResp := response.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    response.LoginFailed,
		}

		var (
			err     error
			request OAuthLoginRequest
		)

		if err = c.ShouldBindJSON(&request); err != nil {
			apiResp.Message = err.Error()
			apiResp.Send(c)
			return
		}

		var (
			token    *oauth2.Token
			userInfo *oauth.UserInfo
		)

		if token, err = provider.GetToken(request.Code); err != nil {
			log.Println(err)
			apiResp.Send(c)
			return
		}
		if userInfo, err = provider.GetUserInfo(token); err != nil {
			log.Println(err)
			apiResp.Send(c)
			return
		}

		// TODO: create or update user data
		helper.PrintJson(userInfo)

		var (
			jwtAuth     auth.JWT
			accessToken string
		)

		if jwtAuth, err = auth.NewJWT(); err != nil {
			log.Println(err)
			apiResp.Send(c)
			return
		}

		// TODO: replace USER_ID to real user id stored in db
		if accessToken, err = jwtAuth.GenerateAccessToken("USER_ID", time.Now().AddDate(0, 3, 0)); err != nil {
			log.Println(err)
			apiResp.Send(c)
			return
		}

		apiResp.StatusCode = http.StatusOK
		apiResp.Result = true
		apiResp.Message = response.Success
		apiResp.Data = OAuthLoginResponse{accessToken}
		apiResp.Send(c)
		return
	}
}
