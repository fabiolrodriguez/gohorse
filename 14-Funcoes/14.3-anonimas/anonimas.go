package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("cebolaZ")

	fmt.Println(retorno)
}
