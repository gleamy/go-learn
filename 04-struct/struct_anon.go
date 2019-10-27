package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	Skills
	int
	speciality string
}

func main() {
	jane := Student{Human: Human{name: "jane", age: 12}, speciality: "Biology"}
	// 初始化时不允许直接赋值
	// j := Student {name: "jane", age:12, int: 1}	

	fmt.Println("Her name is ", jane.name)
	jane.name = "jane 1"
	fmt.Println("Her name is ", jane.name)
	fmt.Println("her age is ", jane.age)
	fmt.Println("Her speciality is ", jane.speciality)

	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
	fmt.Printf("%v", jane.Human)	

}
