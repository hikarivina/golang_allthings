package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {

	john := Person{"John",
		&Address{"123 London", "Lodon", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "jane"
	jane.Address.StreetAddress = "4567 backer St"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

}
