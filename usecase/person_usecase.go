package usecase

import (
	"almoxarifado/entity"
	"almoxarifado/repository"
	"errors"
)

type PersonUsecase struct {
	repository *repository.PersonRepository
}

func NewPersonUsecase(repository *repository.PersonRepository) *PersonUsecase {
	return &PersonUsecase{
		repository: repository,
	}
}

func (u *PersonUsecase) CreateNewPerson(person entity.Person) (entity.Person, error) {

	if person.Name == "" {
		return entity.Person{}, errors.New("Nome precisa ser preenchido")
	}

	if person.Department == "" {
		return entity.Person{}, errors.New("Por favor informe o departamento do Solicitante")
	}

	err := u.repository.CreateNewPerson(&person)

	if err != nil {
		return entity.Person{}, err
	}

	return person, nil
}

func (u *PersonUsecase) DeletePerson(id int, personID int) (bool, error) {
	if id <= 0 {
		return false, errors.New("ID da pessoa precisa ser válido")
	}

	deleted, err := u.repository.DeletePerson(id, personID)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func (u *PersonUsecase) SearchForAllPerson() ([]entity.Person, error) {
	person, err := u.repository.SearchForAllPerson()

	if err != nil {
		return nil, errors.New("Nenhum valor encontrado")
	}

	return person, nil
}

func (u *PersonUsecase) UpdatePerson(id int, person entity.Person) (entity.Person, error) {
	if id <= 0 {
		return entity.Person{}, errors.New("ID da pessoa precisa ser válido")
	}

	if person.Name == "" {
		return entity.Person{}, errors.New("O nome do usuário não pode ser vazio")
	}

	if person.Department == "" {
		return entity.Person{}, errors.New("O usuário precisa possuir um departamento")
	}

	person.ID = uint(id)
	updatedPerson, err := u.repository.UpdatePerson(id, person)
	if err != nil {
		return entity.Person{}, err
	}
	return updatedPerson, nil
}
