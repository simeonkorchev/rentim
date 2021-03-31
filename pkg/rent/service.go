package rent

import (
	"github.com/pkg/errors"
)

type Repository interface {
	AddProperty(property Property) error
	GetProperty(id string) (Property, error)
	GetAllProperties() ([]Property, error)
	DeleteProperty(id string) (Property, error)
}

type PropertyService struct {
	repo Repository
}

func NewPropertyService(repo Repository) PropertyService {
	return PropertyService{repo:repo}
}

func (s PropertyService) SaveProperty(p Property) error {
	if p.location.city == "" || p.location.street == "" {
		return errors.New("invalid property location, city and street can not be empty")
	}

	if p.price.LessThanOrEqualsZero() {
		return errors.New("invalid property price, can not be less than or equal  0")
	}
	return s.repo.AddProperty(p)
}

func (s PropertyService) FindPropertyByID(id string) (Property, error) {
	return s.repo.GetProperty(id)
}

func (s PropertyService) FindAllProperties() ([]Property, error) {
	return s.repo.GetAllProperties()
}

func (s PropertyService) DeletePropertyByID(id string) (Property, error) {
	return s.repo.DeleteProperty(id)
}