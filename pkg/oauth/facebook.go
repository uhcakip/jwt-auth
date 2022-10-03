package oauth

import (
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"io/ioutil"
	"jwt-auth/pkg/helper"
	"net/http"
	"net/url"
	"strings"
)

type facebook struct {
	oauthConfig *oauth2.Config
}

func (fb *facebook) GetConfig() (config *oauth2.Config) {
	return fb.oauthConfig
}

func (fb *facebook) GetToken(code string) (token *oauth2.Token, err error) {
	if token, err = fb.oauthConfig.Exchange(context.TODO(), code); err != nil {
		return
	}

	return
}

func (fb *facebook) GetUserInfo(token *oauth2.Token) (userInfo *UserInfo, err error) {
	var fileds []string
	if fileds, err = helper.GetStructFields(UserInfo{}, "json"); err != nil {
		return
	}

	query := url.Values{}
	query.Add("fields", strings.Join(fileds, ","))
	query.Add("access_token", token.AccessToken)

	var resp *http.Response
	client := fb.oauthConfig.Client(context.TODO(), token)
	api := viper.GetString("facebook.userInfoAPI") + "?" + query.Encode()

	if resp, err = client.Get(api); err != nil {
		return
	}

	var userInfoByte []byte
	userInfo = new(UserInfo)

	if userInfoByte, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	if err = json.Unmarshal(userInfoByte, userInfo); err != nil {
		return

	}

	return
}
