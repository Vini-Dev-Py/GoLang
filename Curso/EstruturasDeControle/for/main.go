package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// Misturando for com ifs

	fmt.Println("")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("O valor %v é par \n", i)
		} else {
			fmt.Printf("O valor %v é impar \n", i)
		}
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
