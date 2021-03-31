package fakes

import (
	"errors"
	"github.com/simeonkorchev/rentim/pkg/rent"
)

type RepositoryFake struct {
	AddPropertyFunc func(property rent.Property) error
	GetPropertyFunc func(id string) (rent.Property, error)
	GetAllPropertiesFunc func() ([]rent.Property, error)
	DeletePropertyFunc func(id string) (rent.Property, error)
}

func (r RepositoryFake) AddProperty(property rent.Property) error {
	if r.AddPropertyFunc == nil {
		return errors.New("AddProperty called but not initialized in this fake")
	}
	return r.AddPropertyFunc(property)
}

func (r RepositoryFake) GetProperty(id string) (rent.Property, error) {
	if r.GetPropertyFunc == nil {
		return rent.Property{}, errors.New("GetProperty called but not initialized in this fake")
	}
	return r.GetPropertyFunc(id)
}

func (r RepositoryFake) GetAllProperties() ([]rent.Property, error) {
	if r.GetAllPropertiesFunc == nil {
		return nil, errors.New("GetAllProperties called but not initialized in this fake")
	}
	return r.GetAllPropertiesFunc()
}

func (r RepositoryFake) DeleteProperty(id string) (rent.Property, error) {
	if r.DeletePropertyFunc == nil {
		return rent.Property{}, errors.New("DeletePropertyByID called but not initialized in this fake")
	}
	return r.DeletePropertyFunc(id)
}