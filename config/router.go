package config

func (pc *PersonController) RegisterRoutes(router gin.IRouter) {
	persons := router.Group("/persons")
	persons.POST("", pc.CreateNewPerson)
	persons.GET("", pc.SearchForAllPerson)
	persons.DELETE("/:id", pc.DeletePerson)
}