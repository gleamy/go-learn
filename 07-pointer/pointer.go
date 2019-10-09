package main

import "fmt"

func main() {
	var ab int = 20 /* 声明实际变量 */
	var ip *int     /* 声明指针变量 */

	ip = &ab        /* 向指针变量中存入地址 */

	fmt.Printf("ab 的值是: %v\n", ab)
	fmt.Printf("ab 的地址是: %v\n", &ab)
	fmt.Printf("ip 存储的地址: %v\n", ip)
	fmt.Printf("ip 的地址: %v\n", &ip)
	fmt.Printf("*ip 的值: %v\n", *ip)
	fmt.Printf("*&ab = %v\n", *&ab)
	fmt.Printf("*&ip = %v\n", *&ip)

	ip = nil

	fmt.Printf("空指针是: %v\n", ip)
	fmt.Printf("空指针的地址是: %v\n", &ip)
	fmt.Printf("空指针的16进制是: %x\n", ip)
}
