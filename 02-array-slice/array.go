package main

import "fmt"

func main() {

	// - 数组的声明
	// var arr1 [5]int
	arr1 := [...]int {1,2,3};
	//   - 值传递
	for i, v := range arr1 {
		arr1[i] = i + 1
		fmt.Printf("arr1 %d, %d\n", i, v)
	}
	for i, v := range arr1 {
		fmt.Printf("arr1 %d, %d\n", i, v)
	}
	// PrintArray("array:", arr1)
	//   - 数组不允许扩展
	// arr1 = append(arr1, 5)


	// - 切片
	arr2 := []int{9, 8, 7, 6}
	PrintArray("slice: ", arr2)
	//   - 可以扩展 
	arr2 = append(arr2, 5)

}

func PrintArray(lab string, arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s %d, %d\n", lab, i, arr[i])
	}

	//for i, v := range arr {
        //        fmt.Printf("%s %d, %d\n", lab, i, v)
        //}
}
