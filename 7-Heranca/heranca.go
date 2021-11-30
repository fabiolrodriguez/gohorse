package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	zeh := estudante{pessoa{"Zé", 67, 154, "da Silva"}, "Engenharia", "Facú"}
	fmt.Println(zeh)

	p1 := pessoa{"João", 20, 220, "Pedrosa"}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(zeh.curso)
}
