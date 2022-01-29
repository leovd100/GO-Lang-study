package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case func():
		return "função"
	case string:
		return "string"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(func() {
	}))
	fmt.Println(tipo(time.Now()))
}
