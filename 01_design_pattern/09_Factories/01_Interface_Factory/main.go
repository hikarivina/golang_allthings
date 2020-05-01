package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Println("Hi all, my name is %s, I am %d years old",
		p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, i'm too tired", p.name, p.age)
}

func NewPerson(name string, age int) Person {

	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("DangNH", 130)
	p.SayHello()
}
