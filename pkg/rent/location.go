package rent

type Location struct {
	city string
	street string
}

func NewLocation(city, street string) Location{
	return Location{
		city:   city,
		street: street,
	}
}