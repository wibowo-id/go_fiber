package wa_sender

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/module/wa_sender/controller"
	"go_fiber_wibowo/app/module/wa_sender/repository"
	"go_fiber_wibowo/app/module/wa_sender/service"
)

// struct of WaSenderRouter
type WaSenderRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of auth module
var NewWaSenderModule = fx.Options(
	// register repository of auth module
	fx.Provide(repository.NewWaSenderRepository),

	// register service of auth module
	fx.Provide(service.NewWaSenderService),

	// register controller of auth module
	fx.Provide(controller.NewController),

	// register router of auth module
	fx.Provide(NewWaSenderRouter),
)

// init WaSenderRouter
func NewWaSenderRouter(fiber *fiber.App, controller *controller.Controller) *WaSenderRouter {
	return &WaSenderRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of auth module
func (_i *WaSenderRouter) RegisterRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.WaSender

	// define routes
	prefix.Route("/send-message", func(router fiber.Router) {
		router.Post("/otp", Controller.SendOtp)
		router.Post("/generate-qr", Controller.GenerateQr)
	})
}
