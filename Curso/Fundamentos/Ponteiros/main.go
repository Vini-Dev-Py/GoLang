package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros

	var ponteiro *int = nil

	ponteiro = &i // pegando o endereço da variável

	fmt.Println(ponteiro)

	*ponteiro++ // soma 1 no valor da variável que o ponteiro esta referenciando

	fmt.Println(ponteiro)

	fmt.Println(*ponteiro)

	fmt.Println(i)
}
