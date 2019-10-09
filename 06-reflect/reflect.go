package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var y float64 = 3.2
	// w := reflect.ValueOf(y)
	p := reflect.ValueOf(&y)
	w := p.Elem()
	// w.SetFloat(5.3)
	w.SetFloat(5.3)
	fmt.Println("p type: ", p.Type())
	fmt.Println("w type: ", w.Type())
	fmt.Println("y is:", y)


	var str string
	fmt.Println("str is: ", reflect.ValueOf(str))
	fmt.Println(str == "") // nil)
}
