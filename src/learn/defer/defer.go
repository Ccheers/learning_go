package main

import "fmt"

//基础defer
//func main() {
//	defer fmt.Println("world")
//
//	fmt.Println("hello")
//}

//defer栈（）
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
