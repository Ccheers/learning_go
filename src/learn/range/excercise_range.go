package main

import "golang.org/x/tour/pic"

//Pic
// dx slice 中每个元素的长度的 8 位无符号整数
// dy slice 的长度
func Pic(dx, dy int) [][]uint8 {
	// 外层slice
	a := make([][]uint8, dy)
	for x := range a {
		// 里层slice
		b := make([]uint8, dx)
		for y := range b {
			// 给里层slice每个元素赋值
			b[y] = uint8(x*y - 1)
		}
		// 给外层slice每个元素赋值
		a[x] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
