package main

import "fmt"

func notaParaConceito(n float64) string {

	notaA := n >= 9
	notaB := n >= 8 && n < 9
	notaC := n >= 5 && n < 8
	notaD := n >= 3 && n < 5

	switch {

	case notaA:
		return "A"
	case notaB:
		return "B"
	case notaC:
		return "C"
	case notaD:
		return "D"
	default:
		return "E"
	}

}

func main() {
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(3))
}
