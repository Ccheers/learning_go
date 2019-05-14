package main

import "fmt"

//type Vertex struct {
//	X int
//	Y float64
//}

//结构体
//func main() {
//	fmt.Println(Vertex{1, 2})
//}
//结构体字段
//func main() {
//	v := Vertex{1, 2}
//	v.Y = 4
//	fmt.Printf("%T",v.Y)
//}
//结构体指针
//func main() {
//	v := Vertex{1, 2}
//	p := &v
//	p.X = 16
//	fmt.Println(v)
//}

//结构体文法
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("%T", p)
}
