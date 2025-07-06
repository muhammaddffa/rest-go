package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"shellrean.id/belajar-golang-rest-api/domain"
	"shellrean.id/belajar-golang-rest-api/dto"
	"shellrean.id/belajar-golang-rest-api/internal/util"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerApi{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error{
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil{
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil{
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}