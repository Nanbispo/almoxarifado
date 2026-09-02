package config

import (
	"almoxarifado/controller"

	"github.com/gin-gonic/gin"
)

func RegisterPersonRoutes(router gin.IRouter, pc *controller.PersonController) {
	persons := router.Group("/persons")
	persons.POST("", pc.CreateNewPerson)
	persons.GET("", pc.SearchForAllPerson)
	persons.DELETE("/:id", pc.DeletePerson)
}
