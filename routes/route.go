package routes

import (
	"fmt"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/token"
	"github.com/adityaputra42/e-commerce_backend/utils"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config     utils.Config
	Store      db.Store
	TokenMaker token.Maker
	Route      *fiber.App
}

func InitServer(config utils.Config, dbstore db.Store) error {

	token, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker %w", err)
	}

	app := fiber.New()
	server := &Server{
		Config:     config,
		Store:      dbstore,
		TokenMaker: token,
		Route:      app,
	}

	server.RouteInit()
	return app.Listen(config.ServerAddress)
}

func (server *Server) RouteInit() {

	server.Route.Group("/api/v1")
	{

	}

}
