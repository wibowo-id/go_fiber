package main

import (
	"dlh_oss_be/app/module/article"
	"dlh_oss_be/app/module/auth"
	"dlh_oss_be/app/module/options"
	"dlh_oss_be/app/module/permission"
	"dlh_oss_be/app/module/role"
	"dlh_oss_be/app/module/role_permission"
	"dlh_oss_be/app/module/user"
	"dlh_oss_be/app/module/version"
	"dlh_oss_be/app/module/wa_sender"
	"dlh_oss_be/app/module/web_page"
	"go.uber.org/fx"

	"dlh_oss_be/app/middleware"
	"dlh_oss_be/app/router"
	"dlh_oss_be/internal/bootstrap"
	"dlh_oss_be/internal/bootstrap/database"
	"dlh_oss_be/utils/config"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
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
		web_page.NewWebPageModule,

		version.NewVersionModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
