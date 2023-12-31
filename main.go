package main

import (
	"go-short-url/internal/config"
	"go-short-url/internal/delivery/http"
	"go-short-url/internal/middleware"
	"go-short-url/internal/repository"
	"go-short-url/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.InitConfig()
	db := config.InitDB()

	urlRepository := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepository)
	urlHandler := http.NewURLHandler(urlService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, urlRepository)
	userHandler := http.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository)
	authHandler := http.NewAuthHandler(authService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "${ip} ${status} - ${method} ${path}\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// auth routes
	v1.Post("/auth/register", authHandler.RegisterUser)
	v1.Post("/auth/login", authHandler.AuthLogin)

	// user routes
	v1.Get("/:user_id", userHandler.GetURLsByUserId) // get list of shortened url by user_id

	// short url routes
	v1.Post("/short_url", middleware.AuthMiddleware, urlHandler.CreateShortURL)
	v1.Get("/short_url/", urlHandler.GetById) // short_url/id?=
	v1.Delete("/short_url/:id", urlHandler.Delete)
	app.Get("/:shortCode", urlHandler.RedirectURL) // redirect

	_ = app.Listen("127.0.0.1:1232")
}
