package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00

	delete(funcsESalarios, "Lucas Souza")

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s recebe um salario de %f\n", nome, salario)
	}
}
