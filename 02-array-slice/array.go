package main

import "fmt"

func main() {

	// - 数组的声明
	// var a1 [5]int
	a1 := [...]int{1, 2, 3}
	for i, v := range a1 {
		fmt.Printf("a1 %d, %d\n", i, v)
	}
	// PrintSlice("a1:", a1)
	//   - 数组不允许扩展
	// a1 = append(a1, 5)

	// - 切片
	s1 := []int{1, 2, 3, 4}
	PrintSlice("s1-0: ", s1)
	//   - 可以扩展
	s1 = append(s1, 5)
	PrintSlice("s1-1: ", s1)

	s2 := s1[2:3]
	PrintSlice("s2-0 s2 := s1[2:3] : ", s2)

	s2 = append(s2, -1);
	PrintSlice("s2-1 s2 = append(s2, -1) : ", s2)
	PrintSlice("s1-2", s1)

	s3 := s1[:2]
	PrintSlice("s3-0 s3 := s1[:2] : ", s3)
}

func PrintSlice(lab string, s []int) {
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%s %d, %d\n", lab, i, arr[i])
	// }
	fmt.Printf("%s len=%d cap=%d slice=%v\n", lab, len(s), cap(s), s)
	for i, v := range s {
		fmt.Printf("%s %d, %d\n", lab, i, v)
	}
}
