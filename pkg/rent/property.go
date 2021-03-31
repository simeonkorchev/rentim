package rent

import "github.com/simeonkorchev/rentim/pkg/money"

type PropertyType int

const (
	Room PropertyType = iota
	OneBedroom
	TwoBedroom
	ThreeBedroom
	FourBedroom
	MultiBedroom
	Penthouse
	House
	Studio
	Office
)

func (p PropertyType) String() string {
	return [...]string{"room", "one bedroom", "two bedrooms", "three bedrooms", "four bedrooms", "multi bedrooms", "penthouse", "house", "studio", "office"}[p]
}

type Property struct {
	id string
	propType PropertyType
	location Location
	size int
	price money.Money
}

func NewProperty(id string, propType PropertyType, location Location, size int, price money.Money) Property {
	return Property{
		id: id,
		propType:propType,
		location: location,
		size: size,
		price: price,
	}
}

func (p Property) Id() string {
	return p.id
}

func (p Property) PropType() PropertyType {
	return p.propType
}

func (p Property) Location() Location {
	return p.location
}

func (p Property) Size() int {
	return p.size
}

func (p Property) Price() money.Money {
	return p.price
}