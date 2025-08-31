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
	cfg config.OauthConfig
}

func New() *OAuthService {
	cfg, _ := env.ParseAs[config.OauthConfig]()
	
	var AppConfig oauth2.Config
	return &OAuthService{AppConfig, cfg}
}

func (s *OAuthService) GoogleConfig() oauth2.Config {
	s.GoogleLoginConfig = oauth2.Config{
		RedirectURL: s.cfg.CLIENT_CALLBACK_URL,
		ClientID: s.cfg.CLIENT_ID,
		ClientSecret: s.cfg.CLIENT_SECRET,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
        Endpoint: google.Endpoint,
	}

	return s.GoogleLoginConfig
}
