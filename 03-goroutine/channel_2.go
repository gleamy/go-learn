package main

import "fmt"

func main() {
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
}
