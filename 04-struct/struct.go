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
	var b person
	p.name = "ming"
	p.age = 10

	fmt.Printf("%v  \n", p)

	fmt.Printf("%v \n", b)

	p1 := new(person)

	p1.name = "m"
	p1.age=22


	fmt.Printf("%v \n", *p1)

}
