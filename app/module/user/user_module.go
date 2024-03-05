package user

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/middleware"
	"go_fiber_wibowo/app/module/user/controller"
	"go_fiber_wibowo/app/module/user/service"
)

// struct of UserRouter
type UserRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of user module
var NewUserModule = fx.Options(
	// register repository of user module
	// already provided in auth module
	//fx.Provide(repository.NewUserRepository),

	// register service of user module
	fx.Provide(service.NewUserService),

	// register controller of user module
	fx.Provide(controller.NewController),

	// register router of user module
	fx.Provide(NewUserRouter),
)

// init UserRouter
func NewUserRouter(fiber *fiber.App, controller *controller.Controller) *UserRouter {
	return &UserRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of user module
func (_i *UserRouter) RegisterUserRoutes(prefix fiber.Router) {
	// define controllers
	Controller := _i.Controller.User

	// define routes
	prefix.Route("/admin/user", func(router fiber.Router) {
		router.Use(middleware.Protected())
		router.Get("/", Controller.Index)
		router.Post("/", Controller.Store)
		router.Get("/:id", Controller.Show)
		router.Put("/:id", Controller.Update)
		router.Delete("/:id", Controller.Delete)
	})
}
