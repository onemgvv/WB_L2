package pkg

type Address struct {
	city   string
	street string
	home   int
	zip    int
}

func NewAddress(city, street string, home, zip int) *Address {
	return &Address{city, street, home, zip}
}
