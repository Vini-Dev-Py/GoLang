package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)

	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Person struct {
		Nome string
	}

	p1 := Person{"Vinicius"}
	p2 := Person{"Isabella"}

	fmt.Println("Pessoas:", p1 == p2)
}
