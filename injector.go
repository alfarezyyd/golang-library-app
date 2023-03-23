//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"golang-library-app/app"
	"golang-library-app/controller"
	controllerImpl "golang-library-app/controller/impl"
	"golang-library-app/repository"
	repositoryImpl "golang-library-app/repository/impl"
	"golang-library-app/usecase"
	usecaseImpl "golang-library-app/usecase/impl"
	"gorm.io/gorm"
)

var bookSet = wire.NewSet(
	repositoryImpl.NewBookRepositoryImpl,
	wire.Bind(new(repository.BookRepository), new(*repositoryImpl.BookRepositoryImpl)),
	usecaseImpl.NewBookUsecaseImpl,
	wire.Bind(new(usecase.BookUsecase), new(*usecaseImpl.BookUsecaseImpl)),
	controllerImpl.NewBookControllerImpl,
	wire.Bind(new(controller.BookController), new(*controllerImpl.BookControllerImpl)),
)

var publisherSet = wire.NewSet(
	repositoryImpl.NewPublisherRepositoryImpl,
	wire.Bind(new(repository.PublisherRepository), new(*repositoryImpl.PublisherRepositoryImpl)),
	usecaseImpl.NewPublisherUsecaseImpl,
	wire.Bind(new(usecase.PublisherUsecase), new(*usecaseImpl.PublisherUsecaseImpl)),
	controllerImpl.NewPublisherControllerImpl,
	wire.Bind(new(controller.PublisherController), new(*controllerImpl.PublisherControllerImpl)),
)

var kindSet = wire.NewSet(
	repositoryImpl.NewKindRepositoryImpl,
	wire.Bind(new(repository.KindRepository), new(*repositoryImpl.KindRepositoryImpl)),
	usecaseImpl.NewKindUsecaseImpl,
	wire.Bind(new(usecase.KindUsecase), new(*usecaseImpl.KindUsecaseImpl)),
	controllerImpl.NewKindControllerImpl,
	wire.Bind(new(controller.KindController), new(*controllerImpl.KindControllerImpl)),
)

var allSet = wire.NewSet(
	bookSet, publisherSet, kindSet,
)

func InitializedGinEngine(databaseSetup *gorm.DB) *gin.Engine {
	wire.Build(app.NewRouter, allSet)
	return nil
}
