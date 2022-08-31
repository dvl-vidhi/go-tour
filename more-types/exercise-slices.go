package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		slice := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slice[j] = uint8((i + j) / 2)
		}
		ans[i] = slice
	}
	return ans
}

func main() {
	pic.Show(Pic)
}
