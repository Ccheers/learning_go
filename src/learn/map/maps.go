package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//map
//map 映射键到值。
//
//map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。

//var m map[string]Vertex

//map 的文法
//map 的文法跟结构体文法相似，不过必须有键名。

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	value, exist := m["Answer"] //value,是否存在(bool)
	fmt.Println("The value:", value, "Present?", exist)
	fmt.Println(m)
}
