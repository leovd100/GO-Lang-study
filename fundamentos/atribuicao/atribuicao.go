package main

import "fmt"

func main() {
	var b byte = 3 // uint8
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	x, y := 1, 2
	x, y = y, x

	fmt.Println(x, y)
}
