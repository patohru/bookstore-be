package oauth

import (
	"bookstore-be/internal/config"

	"github.com/caarlos0/env/v11"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	_ "github.com/joho/godotenv/autoload"
)

type OAuthService struct {
    GoogleLoginConfig oauth2.Config
}

func New() *OAuthService {
	var AppConfig oauth2.Config
	return &OAuthService{AppConfig}
}

func (s *OAuthService) GoogleConfig() oauth2.Config {
	cfg, _ := env.ParseAs[config.OauthConfig]()
	s.GoogleLoginConfig = oauth2.Config{
		RedirectURL: cfg.CLIENT_CALLBACK_URL,
		ClientID: cfg.CLIENT_ID,
		ClientSecret: cfg.CLIENT_SECRET,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return s.GoogleLoginConfig
}
