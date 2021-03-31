package inmem

import (
	"fmt"
	"github.com/simeonkorchev/rentim/pkg/rent"
	"math/rand"
)

type InMemory struct {
	properties map[string]rent.Property
}

func (i InMemory) AddProperty(property rent.Property) (rent.Property, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return rent.Property{}, fmt.Errorf("could not generate id: '%w'", err)
	}

	id := string(b)
	added := rent.NewProperty(id, property.PropType(), property.Location(), property.Size(), property.Price())
	i.properties[id] = added
	return added, nil
}

func (i InMemory) GetProperty(id string) (rent.Property, error) {
	found, ok := i.properties[id]
	if !ok {
		return rent.Property{}, fmt.Errorf("could not found proprety with id '%v'", id)
	}
	return found, nil
}

func (i InMemory) GetAllProperties() ([]rent.Property, error) {
	properties := make([]rent.Property, 0)
	for _, property := range i.properties {
		properties = append(properties, property)
	}
	return properties, nil
}

func (i InMemory) DeleteProperty(id string) (rent.Property, error) {
	prop, ok := i.properties[id]
	if !ok {
		return rent.Property{}, fmt.Errorf("could not found proprety with id '%v'", id)
	}
	delete(i.properties, id)
	return prop, nil
}