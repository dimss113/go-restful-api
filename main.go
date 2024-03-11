package main

import (
	"dimasfadilah/go-restful-api/app"
	"dimasfadilah/go-restful-api/controller"
	"dimasfadilah/go-restful-api/helper"
	"dimasfadilah/go-restful-api/middleware"
	"dimasfadilah/go-restful-api/repository"
	"dimasfadilah/go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
