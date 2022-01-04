package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(10, 20, 30, 1232131, 543995439, 2391932193292391)
	fmt.Println(totalSoma)
	escrever("cebola", 10, 20, 30)
}
