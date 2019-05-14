package main

import "fmt"

//while => for
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//基础for
//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}

//for续
//func main() {
//	sum := 1
//	for ; sum < 1000; {
//		sum += sum
//	}
//	fmt.Println(sum)
//}

//死循环
//func main() {
//	sum := 1
//	for {
//		sum += 1
//	}
//	fmt.Println(sum)
//}
