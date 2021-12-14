package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Numero Inválido"
	}
}
func diaDaSemana2(numero int) string {
	var diaDaSemanaz string
	switch {
	case numero == 1:
		diaDaSemanaz = "Domingo"
	case numero == 2:
		diaDaSemanaz = "Segunda-feira"
	case numero == 3:
		diaDaSemanaz = "Terça-feira"
		fallthrough
	case numero == 4:
		diaDaSemanaz = "Quarta-feira"
	case numero == 5:
		diaDaSemanaz = "Quinta-feira"
	case numero == 6:
		diaDaSemanaz = "Sexta-feira"
	case numero == 7:
		diaDaSemanaz = "Sabado"
	default:
		diaDaSemanaz = "Domingo"
	}

	return diaDaSemanaz
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)

	dia2 := diaDaSemana2(3)

	fmt.Println(dia)

	fmt.Println(dia2)

}
