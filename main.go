package main

import (
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"shellrean.id/belajar-golang-rest-api/dto"
	"shellrean.id/belajar-golang-rest-api/internal/api"
	"shellrean.id/belajar-golang-rest-api/internal/config"
	"shellrean.id/belajar-golang-rest-api/internal/connection"
	"shellrean.id/belajar-golang-rest-api/internal/repository"
	"shellrean.id/belajar-golang-rest-api/internal/service"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("endpoint perlu token, silahkan login dulu"))
		},
	})

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	authRepository := repository.NewUser(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, authRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
