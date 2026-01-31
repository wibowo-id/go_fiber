package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go_fiber_wibowo/app/middleware"
	"go_fiber_wibowo/app/module/article"
	"go_fiber_wibowo/app/module/auth"
	"go_fiber_wibowo/app/module/options"
	"go_fiber_wibowo/app/module/permission"
	"go_fiber_wibowo/app/module/role"
	"go_fiber_wibowo/app/module/role_permission"
	"go_fiber_wibowo/app/module/user"
	"go_fiber_wibowo/app/module/version"
	versionController "go_fiber_wibowo/app/module/version/controller"
	"go_fiber_wibowo/app/module/wa_sender"
	"go_fiber_wibowo/utils/config"
)

type Router struct {
	App fiber.Router
	Cfg *config.Config

	// router
	ArticleRouter        *article.ArticleRouter
	AuthRouter           *auth.AuthRouter
	OptionRouter         *options.OptionRouter
	RoleRouter           *role.RoleRouter
	PermissionRouter     *permission.PermissionRouter
	RolePermissionRouter *role_permission.RolePermissionRouter
	UserRouter           *user.UserRouter
	VersionRouter        *version.VersionRouter
	WaSenderRouter       *wa_sender.WaSenderRouter

	// controller
	VersionController *versionController.Controller
}

func NewRouter(
	fiber *fiber.App,
	cfg *config.Config,

	// router
	articleRouter *article.ArticleRouter,
	authRouter *auth.AuthRouter,
	optionRouter *options.OptionRouter,
	roleRouter *role.RoleRouter,
	permissionRouter *permission.PermissionRouter,
	rolePermissionRouter *role_permission.RolePermissionRouter,
	userRouter *user.UserRouter,
	versionRouter *version.VersionRouter,
	waSenderRouter *wa_sender.WaSenderRouter,

	// controller
	versionController *versionController.Controller,
) *Router {
	return &Router{
		App: fiber,
		Cfg: cfg,

		// router
		ArticleRouter:        articleRouter,
		AuthRouter:           authRouter,
		OptionRouter:         optionRouter,
		RoleRouter:           roleRouter,
		PermissionRouter:     permissionRouter,
		RolePermissionRouter: rolePermissionRouter,
		UserRouter:           userRouter,
		VersionRouter:        versionRouter,
		WaSenderRouter:       waSenderRouter,

		// controller
		VersionController: versionController,
	}
}

// Register routes
func (r *Router) Register() {
	r.App.Use(logger.New(logger.Config{
		Format:     "${time} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	r.App.Use(cors.New())
	// Swagger Documentation
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	api := r.App.Group("/api")

	// Test Routes
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	api.Get("version", r.VersionController.Version.IndexOne)

	// Register routes of modules
	r.AuthRouter.RegisterAuthRoutes(api)
	r.WaSenderRouter.RegisterRoutes(api)

	api.Use(middleware.Protected())
	r.OptionRouter.RegisterOptionRoutes(api)
	r.RoleRouter.RegisterRoleRoutes(api)
	r.PermissionRouter.RegisterPermissionRoutes(api)
	r.RolePermissionRouter.RegisterRolePermissionRoutes(api)
	r.UserRouter.RegisterUserRoutes(api)
	r.ArticleRouter.RegisterArticleRoutes(api)
	r.VersionRouter.RegisterVersionRoutes(api)
}
