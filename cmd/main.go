package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"go_fiber_wibowo/app/middleware"
	"go_fiber_wibowo/app/module/article"
	"go_fiber_wibowo/app/module/auth"
	"go_fiber_wibowo/app/module/options"
	"go_fiber_wibowo/app/module/permission"
	"go_fiber_wibowo/app/module/role"
	"go_fiber_wibowo/app/module/role_permission"
	"go_fiber_wibowo/app/module/user"
	"go_fiber_wibowo/app/module/version"
	"go_fiber_wibowo/app/module/wa_sender"
	"go_fiber_wibowo/app/router"
	"go_fiber_wibowo/internal/bootstrap"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
)

// @title                       Go Fiber Starter API Documentation
// @version                     1.0
// @description                 This is a sample API documentation.
// @termsOfService              http://swagger.io/terms/
// @contact.name                Developer
// @contact.email               bangadam.dev@gmail.com
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host                        localhost:8080
// @schemes                     http https
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 "Type 'Bearer {TOKEN}' to correctly set the API Key"
// @BasePath                    /
func main() {
	fx.New(

		/* provide patterns */
		// config
		fx.Provide(config.NewConfig),
		// logging
		fx.Provide(bootstrap.NewLogger),
		// fiber
		fx.Provide(bootstrap.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules
		article.NewArticleModule,
		auth.NewAuthModule,
		options.NewOptionModule,
		wa_sender.NewWaSenderModule,
		role.NewRoleModule,
		permission.NewPermissionModule,
		user.NewUserModule,
		role_permission.NewRolePermissionModule,

		version.NewVersionModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
