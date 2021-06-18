package repository

type Repository interface {
	Save(data interface{}) error
	GetAll(params ...interface{}) (interface{}, error)
	Get(params ...interface{}) (interface{}, error)
	Delete(param interface{}) error
}
