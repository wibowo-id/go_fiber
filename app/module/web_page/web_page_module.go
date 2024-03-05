package web_page

import (
	"dlh_oss_be/app/module/web_page/controller"
	"dlh_oss_be/app/module/web_page/repository"
	"dlh_oss_be/app/module/web_page/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of WebPageRouter
type WebPageRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of webPage module
var NewWebPageModule = fx.Options(
	// register repository of webPage module
	fx.Provide(repository.NewWebPageRepository),

	// register service of webPage module
	fx.Provide(service.NewWebPageService),

	// register controller of webPage module
	fx.Provide(controller.NewController),

	// register router of webPage module
	fx.Provide(NewWebPageRouter),
)

// init WebPageRouter
func NewWebPageRouter(fiber *fiber.App, controller *controller.Controller) *WebPageRouter {
	return &WebPageRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of webPage module
func (_i *WebPageRouter) RegisterWebPageRoutes(prefix fiber.Router) {
	// define controllers
	WebPageController := _i.Controller.WebPage

	// define routes
	prefix.Route("/admin/web-page", func(router fiber.Router) {
		router.Get("/", WebPageController.Index)
		router.Get("/:id", WebPageController.Show)
		router.Post("/", WebPageController.Store)
		router.Put("/:id", WebPageController.Update)
		router.Delete("/:id", WebPageController.Delete)
	})
}
