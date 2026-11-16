//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-dependency-injection/controller"
	"belajar-golang-dependency-injection/middleware"
	"belajar-golang-dependency-injection/repository"
	"belajar-golang-dependency-injection/service"
	"net/http"

	"belajar-golang-dependency-injection/app"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)
func NewValidator() *validator.Validate {
	return validator.New()
}

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
	NewValidator,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
