package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando Usuário", u.nome)
}

func (u usuario) maior18() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Cebola Roxa", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"David Beckham", 100}
	usuario2.salvar()
	idadeusuario2 := usuario2.maior18()
	fmt.Println(idadeusuario2)
	usuario2.aniversario()
	fmt.Println(usuario2.idade)
}
