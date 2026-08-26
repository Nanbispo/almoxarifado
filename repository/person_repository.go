package repository

import (
	"almoxarifado/entity"
	"database/sql"
)

type PersonRepository struct {
	database *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{
		database: database,
	}
}

func (r *PersonRepository) CreateNewPerson(person *entity.Person) error {
	query := `
		INSERT INTO person(username, department) 
		VALUES ($1, $2)
		RETURNING id
	`

	return r.database.QueryRow(
		query,
		person.Name,
		person.Department,
	).Scan(&person.ID)

}

func (r *PersonRepository) DeletePerson(id int, personID int) (bool, error) {

	result, err := r.database.Exec("UPDATE person SET D_E_L_E_T_ = '*' WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil

}
