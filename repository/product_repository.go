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

func (r *ProductRepository) SearchForAllProduct() ([]entity.Product, error) {
	query := `
		SELECT id, name, description, department
		FROM PRODUCT
		WHERE (D_E_L_E_T_ = '' OR D_E_L_E_T_ IS NULL)
		AND DATBLO IS NULL
	`

	rows, err := r.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var searchProduct []entity.Product

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Department,
		); err != nil {
			return nil, err
		}

		searchProduct = append(searchProduct, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return searchProduct, nil
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
