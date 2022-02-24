package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[87987654321] = "Pedro"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])

	delete(aprovados, 12345678978)

	fmt.Println(aprovados)
}
