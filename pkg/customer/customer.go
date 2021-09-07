package customer

type Customer struct {
	name    string
	address string
	phone   string
}

//new return a new costumer

func new(name string, address string, phone string) Customer {
	return Customer{name, address, phone}
}
