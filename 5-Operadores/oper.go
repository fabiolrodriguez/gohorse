package main

import "fmt"

func main() {
	//Aritiméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multi := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multi, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25

	somaz := numero1 + numero2

	fmt.Println(somaz)
	// fim Aritméticos

	// Atribuição
	var variavel string = "string"
	variavel1 := "string2"

	fmt.Println(variavel, variavel1)
	//fim atribuição

	//Relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Fim relacionais

	//Logicos
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// Fim dos lógicos

	//Unários

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--

	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	// fim unários

	var texto string

	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}

	fmt.Println(texto)

}
