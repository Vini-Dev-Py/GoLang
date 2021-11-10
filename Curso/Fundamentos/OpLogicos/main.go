package main

import "fmt"

func compras(TrabalhoTerca bool, TrabalhoQuinta bool) (bool, bool, bool) {
	comprarTv50 := TrabalhoTerca && TrabalhoQuinta
	comprarTv32 := TrabalhoTerca != TrabalhoQuinta
	comprarSorvete := TrabalhoTerca || TrabalhoQuinta

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
