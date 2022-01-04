package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada, retornando resultado")
	fmt.Println("Função aluno aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	defer funcao1()
	funcao2()
	fmt.Println(alunoAprovado(6, 7))
}
