package v1

import (
	"github.com/TatarinAlba/WBTest/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OrderController struct {
	service *usecase.OrderUsecase
}

func NewOrderController(service *usecase.OrderUsecase) *OrderController {
	return &OrderController{service}
}

func (controller *OrderController) Route(app *fiber.App) {
	app.Get("/v1/api/order/:uid", controller.getOrder)
	app.Get("/", controller.indexPage)
}

func (controller *OrderController) indexPage(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}

func (controller *OrderController) getOrder(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")
	order, err := controller.service.GetOrder(uid)
	if err != nil {
		logrus.Errorf("Cannot get order with uid [%s]: %s", uid, err)
		return ctx.Status(400).JSON(struct {
			Status int64  `json:"status"`
			Msg    string `json:"msg"`
		}{400, "Cannot get order by given uid"})
	}
	return ctx.Status(200).JSON(order)
}
