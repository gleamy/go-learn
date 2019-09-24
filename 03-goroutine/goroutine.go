package main

import (
	"fmt"
	"runtime"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}

func main() {
	go say("world !!!")
	say("Hello ")
}
