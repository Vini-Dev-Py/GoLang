package main

func notaParaConceito(n float64) string {

	var nota = int(n)

	switch nota {

	case 10:
		fallthrough // executa o proximo bloco
	case 9:
		return "A"
	case 8, 7: // forma mais elegante de realizar com duas condiciones
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: // se nenhum dos casos acontecer
		return "Nota Inválida"
	}

}

func main() { // comentario teste
	notaParaConceito(10)
	notaParaConceito(9)
	notaParaConceito(8)
	notaParaConceito(7)
	notaParaConceito(6)
	notaParaConceito(5)
	notaParaConceito(4)
	notaParaConceito(3)
	notaParaConceito(2)
	notaParaConceito(1)
	notaParaConceito(0)
	notaParaConceito(1000)
}
