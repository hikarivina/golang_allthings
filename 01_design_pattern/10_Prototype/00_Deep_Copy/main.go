package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {

	john := Person{"John", &Address{"123 London", "Lodon", "UK"}}

	jane := john
	jane.Name = "Jane"
	// Shallow sopy
	// jane.Address.StreetAddress = "456 Baker St"
	jane.Address = &Address{
		"456 backer St",
		john.Address.City,
		john.Address.Country}

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

}
