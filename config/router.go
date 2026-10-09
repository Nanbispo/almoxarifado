package config

import (
	"almoxarifado/controller"

	"github.com/gin-gonic/gin"
)

func RegisterPersonRoutes(router gin.IRouter, pc *controller.PersonController) {
	persons := router.Group("/persons")
	persons.POST("", pc.CreateNewPerson)
	persons.GET("", pc.SearchForAllPerson)
	persons.PUT("/:id", pc.UpdatePersonController)
	persons.DELETE("/:id", pc.DeletePerson)
}

func RegisterProductRoutes(router gin.IRouter, pc *controller.ProductController) {
	products := router.Group("/products")
	products.GET("", pc.SearchForAllProduct)
	router.GET("/product", pc.SearchForAllProduct)
}
