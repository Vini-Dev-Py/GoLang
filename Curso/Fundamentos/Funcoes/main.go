package main

import "fmt"

func Somar(x float64, y float64) float64 {
	return x + y
}

func Imprimir(x float64) {
	fmt.Println(x)
}

func main() {

	result := Somar(5, 10)

	Imprimir(result)

}
