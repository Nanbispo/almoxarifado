package config

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"almoxarifado/controller"
	"almoxarifado/entity"

	"github.com/gin-gonic/gin"
)

type productServiceStub struct{}

func (productServiceStub) SearchForAllProduct() ([]entity.Product, error) {
	return []entity.Product{}, nil
}

func TestRegisterProductRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	productController := controller.NewProductController(productServiceStub{})
	RegisterProductRoutes(router, productController)

	for _, path := range []string{"/products", "/product"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("GET %s returned status %d, want %d", path, response.Code, http.StatusOK)
			}
		})
	}
}
