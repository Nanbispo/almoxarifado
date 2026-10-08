package usecase

import (
	"almoxarifado/entity"
	"almoxarifado/repository"
	"errors"
)

type ProductUsecase struct {
	repository *repository.ProductRepository
}

func NewProductUsecase(productRepository *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repository: productRepository}
}

func (u *ProductUsecase) CreateNewProduct(product entity.Product) (entity.Product, error) {
	if product.Name == "" {
		return entity.Product{}, errors.New("Nome do produto precisa ser preenchido")
	}
	if product.Department == "" {
		return entity.Product{}, errors.New("Departamento precisa ser preenchido")
	}
	if err := u.repository.CreateNewProduct(&product); err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (u *ProductUsecase) DeleteProduct(id int) (bool, error) {
	if id <= 0 {
		return false, errors.New("ID do produto precisa ser válido")
	}
	return u.repository.DeleteProduct(id)
}
