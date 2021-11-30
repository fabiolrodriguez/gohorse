package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Fabio de Lima"
	u.idade = 40
	fmt.Println(u)

	enderecoEx := endereco{"Rua dos bobos", 0}

	u2 := usuario{"Fabio de Loma", 5, enderecoEx}
	fmt.Println(u2)

	u3 := usuario{nome: "Ze"}
	fmt.Println(u3)

}
