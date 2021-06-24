package services

import (
	"user/entity"
	"user/repository"
)

type userServices struct{}

var (
	repo repository.Repository = repository.NewMongoRepository()
)

func (*userServices) Create(data interface{}) error {
	dataInserted := data.(entity.CustomerUser)

	if err := repo.Save(dataInserted); err != nil {
		return err
	}

	return nil
}
func (*userServices) Find(params ...interface{}) (interface{}, error) {
	data, err := repo.Get(params...)

	if err != nil {
		return nil, err
	}
	return data.(entity.CustomerUser), nil
}
func (*userServices) FindAll(params ...interface{}) (interface{}, error) {
	data, err := repo.GetAll()

	if err != nil {
		return nil, err
	}
	return data.([]*entity.CustomerUserResponses), nil
}

func (*userServices) Update(params ...interface{}) error {
	id := params[0].(string)

	if err := repo.Update(id, params[1]); err != nil {
		return err
	}

	return nil
}

func (*userServices) Delete(params ...interface{}) error {
	id := params[0].(string)

	if err := repo.Delete(id); err != nil {
		return err
	}

	return nil
}
func NewServices() Services {
	return &userServices{}
}
