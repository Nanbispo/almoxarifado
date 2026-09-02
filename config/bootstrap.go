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

	router := gin.Default()
	RegisterPersonRoutes(router, personController)

	return router, nil
}
