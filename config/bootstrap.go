package config

import (
	"almoxarifado/controller"
	"almoxarifado/repository"
	"almoxarifado/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeApp(database *sql.DB) (*gin.Engine, error) {
	personRepo := repository.NewPersonRepository(database)
	personUsecase := usecase.NewPersonUsecase(personRepo)
	personController := controller.NewPersonController(personUsecase)

	productRepo := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productController := controller.NewProductController(productUsecase)

	router := gin.Default()
	RegisterPersonRoutes(router, personController)
	RegisterProductRoutes(router, productController)

	return router, nil
}
