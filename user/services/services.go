package services

type Services interface {
	Create(data interface{}) error
	Find(params ...interface{}) (interface{}, error)
	FindAll(params ...interface{}) (interface{}, error)
	Update(params ...interface{}) error
}
