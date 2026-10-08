package repository

import (
	"almoxarifado/entity"
	"database/sql"
)

type ProductRepository struct {
	database *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{database: database}
}

func (r *ProductRepository) CreateNewProduct(product *entity.Product) error {

	query := `
		INSERT INTO product(name, description, department) VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.database.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Department,
	).Scan(&product.ID)
}

func (r *ProductRepository) DeleteProduct(id int) (bool, error) {

	result, err := r.database.Exec("UPDATE product SET D_E_L_E_T_ = '*', DATBLO = CURRENT_DATE WHERE id = $1", id)

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}
	return rows > 0, nil
}
