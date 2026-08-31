package controller

import (
	"almoxarifado/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PersonService contains the person operations used by the HTTP layer.
// Keeping this dependency as an interface makes the controller independent of
// the persistence implementation and straightforward to test.
type PersonService interface {
	CreateNewPerson(entity.Person) (entity.Person, error)
	DeletePerson(id int, personID int) (bool, error)
	SearchForAllPerson() ([]entity.Person, error)
}

type PersonController struct {
	personUsecase PersonService
}

func NewPersonController(personUsecase PersonService) *PersonController {
	return &PersonController{personUsecase: personUsecase}
}

// RegisterRoutes registers the person resource endpoints under /persons.


// CreateNewPerson handles POST /persons.
func (pc *PersonController) CreateNewPerson(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPerson, err := pc.personUsecase.CreateNewPerson(person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPerson)
}

// SearchForAllPerson handles GET /persons.
func (pc *PersonController) SearchForAllPerson(c *gin.Context) {
	persons, err := pc.personUsecase.SearchForAllPerson()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

// DeletePerson handles DELETE /persons/:id. The optional person_id query parameter
// is forwarded for compatibility with the current use-case signature.
func (pc *PersonController) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id deve ser um número positivo"})
		return
	}

	personID := 0
	if value := c.Query("person_id"); value != "" {
		personID, err = strconv.Atoi(value)
		if err != nil || personID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "person_id deve ser um número positivo"})
			return
		}
	}

	deleted, err := pc.personUsecase.DeletePerson(id, personID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "pessoa não encontrada"})
		return
	}

	c.Status(http.StatusNoContent)
}
