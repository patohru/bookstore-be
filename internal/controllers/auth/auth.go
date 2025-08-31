package auth

import (
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgxpool"

	"bookstore-be/internal/database"
	"bookstore-be/internal/services/jwt"
	"bookstore-be/internal/services/oauth"
)

type AuthRoutes struct {
	db				*pgxpool.Pool
	jwtService		*jwt.JwtService
	oauthService	*oauth.OAuthService
}

func RegisterRoutes(f *fuego.Server) {
	r := AuthRoutes{
		db:				database.NewPool(),
		jwtService:		jwt.New(),
		oauthService:	oauth.New(),
	}

	r.oauthService.GoogleConfig()

	fuego.Post(f, "/auth/login", r.LoginHandler)
	fuego.Post(f, "/auth/register", r.RegisterHandler)
	// fuego.GetStd(f, "/auth/google/callback", r.GetAuthCallback)
	fuego.Get(f, "/auth/google", r.GoogleLogin)
}
