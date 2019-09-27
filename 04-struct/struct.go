package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var p person
	//p := new(person)

	p.name = "ming"
	p.age = 10

	fmt.Printf("%v  \n", p)
}
