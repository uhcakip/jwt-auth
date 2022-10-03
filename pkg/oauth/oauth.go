package oauth

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	oauth2Facebook "golang.org/x/oauth2/facebook"
)

type Provider interface {
	GetConfig() *oauth2.Config
	GetToken(code string) (token *oauth2.Token, err error)
	GetUserInfo(token *oauth2.Token) (userInfo *UserInfo, err error)
}

type ProviderName string

const (
	ProviderNameFacebook ProviderName = "facebook"
	// ServiceGoogle   ProviderName = "google"
)

type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewProvider(name ProviderName) Provider {
	domain := viper.GetString("app.domain")

	switch name {
	case ProviderNameFacebook:
		return &facebook{
			oauthConfig: &oauth2.Config{
				ClientID:     viper.GetString("facebook.clientID"),
				ClientSecret: viper.GetString("facebook.clientSecret"),
				RedirectURL:  domain + viper.GetString("facebook.redirectURL"),
				Endpoint:     oauth2Facebook.Endpoint,
				Scopes:       []string{"email", "public_profile"},
			},
		}
	}

	return nil
}
