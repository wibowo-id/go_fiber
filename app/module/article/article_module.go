package article

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/module/article/controller"
	"go_fiber_wibowo/app/module/article/repository"
	"go_fiber_wibowo/app/module/article/service"
)

// struct of ArticleRouter
type ArticleRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of article module
var NewArticleModule = fx.Options(
	// register repository of article module
	fx.Provide(repository.NewArticleRepository),

	// register service of article module
	fx.Provide(service.NewArticleService),

	// register controller of article module
	fx.Provide(controller.NewController),

	// register router of article module
	fx.Provide(NewArticleRouter),
)

// init ArticleRouter
func NewArticleRouter(fiber *fiber.App, controller *controller.Controller) *ArticleRouter {
	return &ArticleRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of article module
func (_i *ArticleRouter) RegisterArticleRoutes(prefix fiber.Router) {
	// define controllers
	articleController := _i.Controller.Article

	// define routes
	prefix.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Put("/:id", articleController.Update)
		router.Delete("/:id", articleController.Delete)
	})
}
