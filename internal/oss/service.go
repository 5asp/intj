package oss

type svc struct{}

func NewService() Service {
	return &svc{}
}

type Service interface {
	add(name string) error
	remove(id int) error
	// getAll() ([]model, error)
}
