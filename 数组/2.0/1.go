package main

import (
	"fmt"
)

// 数组问题，在函数中传递的时候是传递之数值，如果数组很大，对内存的占用会很大，所以在函数中传递数组的地址是更好的选择。
func main() {

	fmt.Println("Hello, world!")
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr)
	fmt.Println(arr)

	modifyArr1(&arr)
	fmt.Println(arr)
}
func modifyArr(a [5]int) {
	a[1] = 20
}

func modifyArr1(a *[5]int) {
	a[1] = 20
}

//
