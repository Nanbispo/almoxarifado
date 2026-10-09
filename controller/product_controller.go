package controller

import (
	"almoxarifado/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	SearchForAllProduct() ([]entity.Product, error)
}

type ProductController struct {
	productUsecase ProductService
}

func NewProductController(productUsecase ProductService) *ProductController {
	return &ProductController{productUsecase: productUsecase}
}

// SearchForAllProduct handles GET /products.
func (pc *ProductController) SearchForAllProduct(c *gin.Context) {
	products, err := pc.productUsecase.SearchForAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
