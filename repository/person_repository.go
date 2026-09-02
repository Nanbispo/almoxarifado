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
		INSERT INTO person(name, department) 
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

	result, err := r.database.Exec("UPDATE person SET D_E_L_E_T_ = '*', DATBLO = CURRENT_DATE WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (r *PersonRepository) SearchForAllPerson() ([]entity.Person, error) {

	query := `SELECT id, name, department 
				FROM person 
				WHERE (D_E_L_E_T_ = '' OR D_E_L_E_T_ IS NULL)
  				AND DATBLO IS NULL;`

	rows, err := r.database.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var searchPerson []entity.Person

	for rows.Next() {
		var person entity.Person

		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Department,
		)

		if err != nil {
			return nil, err
		}

		searchPerson = append(searchPerson, person)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return searchPerson, nil
}
