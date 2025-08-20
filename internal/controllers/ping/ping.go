package ping

import "github.com/go-fuego/fuego"

type PingResponse struct {
	Message string `json:"message"`
}

func PingHandler(c fuego.ContextNoBody) (*PingResponse, error) {
	return &PingResponse{
		Message: "Pong",
	}, nil
}

func RegisterRoutes(f *fuego.Server) {
	fuego.Get(f, "/", PingHandler)
}
