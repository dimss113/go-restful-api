//go:build wireinject
// +build wireinject

package main

import (
	"dimasfadilah/go-restful-api/app"
	"dimasfadilah/go-restful-api/controller"
	"dimasfadilah/go-restful-api/middleware"
	"dimasfadilah/go-restful-api/repository"
	"dimasfadilah/go-restful-api/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var validateSet = wire.NewSet(
	app.NewValidator,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		categorySet,
		validateSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
