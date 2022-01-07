package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	fmt.Println(numero)
	inverterSinalPonteiro(&numero)
	fmt.Println(numero)
}
