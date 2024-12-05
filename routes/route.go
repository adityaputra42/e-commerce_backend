package routes

import (
	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/utils"
	"github.com/gofiber/fiber"
)

type Server struct {
	config utils.Config
	store  db.Store
}

func InitServer(config utils.Config, store db.Store) error {
	app := fiber.New()
	server := &Server{
		store:  store,
		config: config}
	server.RouteInit(app)
	return app.Listen(config.ServerAddress)
}

func (server *Server) RouteInit(app *fiber.App) {

	app.Group("/api/v1")
	{

	}

}
