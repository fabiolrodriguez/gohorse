package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo o main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("fabio.rodriguez@gmail.com")
	fmt.Println(erro)
}
