package auth

import (
	"context"

	"github.com/go-fuego/fuego"
	"bookstore-be/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/register: user's email and password.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRoutes) RegisterHandler(c fuego.ContextWithBody[database.CreateUserParams]) (string, error) {

	request, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid register data",
			Err: err,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Invalid register data",
			Err: err,
		}
	}
	request.Password = string(hashedPassword)

	queries := database.New(r.db) 
	ctx := context.Background()
	id, err := queries.CreateUser(ctx, request)
	if err != nil {
		return "", fuego.BadRequestError{
			Title: "Account with given email already existed",
			Err: err,
		}
	}

	return id.String(), nil
}
