package main

import "fmt"

func imprimiResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimiResultado(7)
	imprimiResultado(6)
}
