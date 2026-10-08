package repository

import (
	"almoxarifado/entity"
	"database/sql"
	"errors"
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

	result, err := r.database.Exec("UPDATE person SET D_E_L_E_T_ = '*', DATBLO = CURRENT_DATE WHERE id = $1", id)

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

func (r *PersonRepository) UpdatePerson(id int, person entity.Person) (entity.Person, error) {
	query := `UPDATE person
		SET name = $1, department = $2
		WHERE id = $3
		  AND (D_E_L_E_T_ = '' OR D_E_L_E_T_ IS NULL)
		  AND DATBLO IS NULL
		RETURNING id, name, department`

	err := r.database.QueryRow(query, person.Name, person.Department, id).Scan(
		&person.ID, &person.Name, &person.Department,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return entity.Person{}, errors.New("pessoa não encontrada")
	}
	if err != nil {
		return entity.Person{}, err
	}
	return person, nil
}
