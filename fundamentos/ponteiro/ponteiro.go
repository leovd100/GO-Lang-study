package main

import "fmt"

func main() {
	i := 1 // 64 bits ou 8 bytes

	var p *int = nil

	p = &i
	fmt.Println(p, *p, i)

	*p++

	i++
	fmt.Println(p, *p, i)
}
