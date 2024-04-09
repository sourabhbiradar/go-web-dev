package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func main() {
	person := NewPerson("Abc", 20)

	fmt.Println(person.name)
	fmt.Println(person.age)
}
