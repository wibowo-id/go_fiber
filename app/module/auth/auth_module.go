package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/middleware"
	"go_fiber_wibowo/app/module/auth/controller"
	"go_fiber_wibowo/app/module/auth/service"
	user_repo "go_fiber_wibowo/app/module/user/repository"
)

// struct of AuthRouter
type AuthRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of auth module
var NewAuthModule = fx.Options(
	// register repository of auth module
	fx.Provide(user_repo.NewUserRepository),

	// register service of auth module
	fx.Provide(service.NewAuthService),

	// register controller of auth module
	fx.Provide(controller.NewController),

	// register router of auth module
	fx.Provide(NewAuthRouter),
)

// init AuthRouter
func NewAuthRouter(fiber *fiber.App, controller *controller.Controller) *AuthRouter {
	return &AuthRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of auth module
func (_i *AuthRouter) RegisterAuthRoutes(prefix fiber.Router) {
	// define controllers
	authController := _i.Controller.Auth

	// define routes
	prefix.Route("/auth", func(router fiber.Router) {
		router.Post("/admin/login", authController.LoginAdmin)
		router.Post("/login", authController.Login)
		router.Post("/register", authController.Register)
		router.Post("/forgot-password", authController.ForgotPassword)
		router.Post("/change-password-public", authController.ChangePasswordPublic)
		router.Post("/request-otp", authController.RequestOTP)
		router.Post("/verify-otp", authController.VerifyOTP)

		router.Use(middleware.Protected())
		router.Post("/logout", authController.Logout)
		router.Get("/profile", authController.Profile)
		router.Post("/change-password", authController.ChangePassword)
		router.Put("/profile", authController.UpdateProfile)
		router.Put("/firebase-token", authController.UpdateFirebaseToken)

	})

	prefix.Route("/admin", func(router fiber.Router) {
		router.Use(middleware.Protected())
		router.Get("/menu", authController.Menu)
		router.Get("/dashboard", authController.Dashboard)
	})

}
