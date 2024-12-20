package api

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

func InitServer(config utils.Config, dbstore db.Store) (*Server, error) {

	token, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	app := fiber.New()
	server := &Server{
		Config:     config,
		Store:      dbstore,
		TokenMaker: token,
		Route:      app,
	}

	server.RouteInit()

	return server, nil
}

func (server *Server) RouteInit() {
	UserController := NewUserController(*server)

	api := server.Route.Group("/api/v1")

	{
		api.Post("/register", UserController.CreateUser)
		api.Post("/login", UserController.Login)
		api.Post("/admin/register", UserController.CreateAdmin)

	}

}
func (server *Server) Start(address string) error {
	return server.Route.Listen(address)

}
