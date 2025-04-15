package api

import (
	"fmt"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/middleware"
	"github.com/adityaputra42/e-commerce_backend/middleware/role"
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
	User := NewUserController(*server)
	Address := NewAddressController(*server)
	Product := NewProductController(*server)
	Category := NewCategoryController(*server)
	ColorVarian := NewColorVarianController(*server)
	SizeVarian := NewSizeVarianController(*server)
	Shipping := NewShippingController(*server)
	PaymentMethod := NewPaymentMethodController(*server)
	Transaction := NewTransactionsController(*server)
	Order := NewOrderController(*server)
	Payment := NewPaymentController(*server)
	Session := NewSessionController(*server)

	api := server.Route.Group("/api/v1").Use(middleware.LoggerMiddleware)

	{
		api.Post("/register", User.CreateUser)
		api.Post("/login", User.Login)
		api.Post("/admin/register", User.CreateAdmin)
		api.Post("/token/renew_token", Session.RenewSession)
		api.Get("/categories", Category.GetAll)
		api.Get("/categories/:id", Category.GetById)
		api.Get("/Product/:product_id", Product.FetchProduct)
		api.Get("/Product", Product.FetchProduct)
	}

	api_user := api.Group("/users").Use(middleware.AuthMiddleware, role.MemberAuth)
	{
		api_user.Get("/me", User.FetchUser)
		api_user.Put("/change_password", User.UpdatePassword)
		api_user.Post("/address/create", Address.CreateAddress)
		api_user.Get("/address", Address.FetchAllAddressByUser)
		api_user.Get("/address/:address_id", Address.FetchAddress)
		api_user.Delete("/address/delete/:id", Address.Delete)
		api_user.Put("/address/update/:id", Address.Update)

		api_user.Get("/shippings", Shipping.GetAll)
		api_user.Get("/shippings/:id", Shipping.GetById)

		api_user.Get("/payment_method", PaymentMethod.GetAll)
		api_user.Get("/payment_method/:id", PaymentMethod.GetById)

		api_user.Get("/payment", Payment.GetAll)
		api_user.Get("/payment/:id", Payment.GetById)
		api_user.Put("/payment/create", Payment.Create)

		api_user.Get("/transactions", Transaction.GetAll)
		api_user.Get("/transactions/:id", Transaction.GetById)
		api_user.Post("/transactions/create", Transaction.Create)

		api_user.Get("/orders", Order.GetAll)
		api_user.Get("/orders/:id", Order.GetById)

	}

	api_admin := api.Group("/admin").Use(middleware.AuthMiddleware, role.AdminAuth)
	{
		api_admin.Get("/me", User.FetchUser)
		api_admin.Get("/Address", Address.FetchAllAddressFromAdmin)
		api_admin.Get("/users", User.FetchAllUSer)
		api_admin.Delete("/users/delete/:id", User.Delete)

		api_admin.Get("/color_varians/:product_id", ColorVarian.GetALl)
		api_admin.Get("/color_varians/:id", ColorVarian.GetById)
		api_admin.Post("/color_varians", ColorVarian.Create)
		api_admin.Put("/color_varians/:product_id", ColorVarian.Update)
		api_admin.Delete("/color_varians/:id", ColorVarian.Delete)

		api_admin.Get("/size_varians/:color_varian_id", SizeVarian.GetAll)
		api_admin.Get("/size_varians/:id", SizeVarian.GetById)
		api_admin.Post("/size_varians", SizeVarian.Create)
		api_admin.Put("/size_varians/:color_varian_id", SizeVarian.Update)
		api_admin.Delete("/size_varians/:id", SizeVarian.Delete)

		api_admin.Get("/categories", Category.GetAll)
		api_admin.Get("/categories/:id", Category.GetById)
		api_admin.Post("/categories", Category.Create)
		api_admin.Put("/categories/:id", Category.Update)
		api_admin.Delete("/categories/:id", Category.Delete)

		api_admin.Get("/shippings", Shipping.GetAll)
		api_admin.Get("/shippings/:id", Shipping.GetById)
		api_admin.Post("/shippings", Shipping.Create)
		api_admin.Put("/shippings/:id", Shipping.Update)
		api_admin.Delete("/shippings/:id", Shipping.Delete)

		api_admin.Get("/payment_method", PaymentMethod.GetAll)
		api_admin.Get("/payment_method/:id", PaymentMethod.GetById)
		api_admin.Post("/payment_method", PaymentMethod.Create)
		api_admin.Put("/payment_method/:id", PaymentMethod.Update)
		api_admin.Delete("/payment_method/:id", PaymentMethod.Delete)

		api_admin.Get("/payment", Payment.GetAll)
		api_admin.Get("/payment/:id", Payment.GetById)
		api_admin.Put("/payment/:id", Payment.Update)
		api_admin.Delete("/payment/:id", Payment.Delete)

		api_admin.Get("/transactions", Transaction.GetAll)
		api_admin.Get("/transactions/:id", Transaction.GetById)
		api_admin.Put("/transactions/:id", Transaction.Update)
		api_admin.Delete("/transactions/:id", Transaction.Delete)

		api_admin.Get("/orders", Order.GetAll)
		api_admin.Get("/orders/:id", Order.GetById)
		api_admin.Put("/orders/:id", Order.Update)
		api_admin.Delete("/orders/:id", Order.Delete)

	}

}
func (server *Server) Start(address string) error {
	return server.Route.Listen(address)

}
