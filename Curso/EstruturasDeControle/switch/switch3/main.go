package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float64:
		return "real"
	case string:
		return "string"
	case func():
		return "func"
	case bool:
		return "bool"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(3))
	fmt.Println(tipo(2.5))
	fmt.Println(tipo(true))
	fmt.Println(tipo("Hello"))
}
