package repository

import (
	"almoxarifado/entity"
	"database/sql"
)

type ProductRepository struct {
	database *sql.DB
}

func NewProductRepostory(database *sql.DB) *ProductRepository {
	return &ProductRepository{database: database}
}

func (r *PersonRepository) CreateNewProduct(product *entity.Product) error {

	query := `
		Insert into product(name, description, department) VALUES ($1, $2, $3)
	`

	return r.database.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Department,
	).Scan(&product.ID)
}
