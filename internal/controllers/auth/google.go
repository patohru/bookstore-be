package auth

import (
	"github.com/go-fuego/fuego"
)

func (r *AuthRoutes) GoogleLogin(c fuego.ContextNoBody) (any, error) {
	url := r.oauthService.GoogleLoginConfig.AuthCodeURL("randomstate")
	
	c.Redirect(301, url)
	return url, nil
}
