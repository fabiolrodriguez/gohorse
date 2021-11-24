package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 10000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//rune = int32
	var numero3 rune = 12345
	fmt.Println(numero3)

	//Numeros reais

	var numeroReal1 float32 = 123.32
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12323232323.32
	fmt.Println(numeroReal2)

	numeroReal3 := 12312.22
	fmt.Println(numeroReal3)
	//Fim numeros reais

	//Strings

	var str string = "Ferro"
	fmt.Println(str)

	srt2 := "Ferro2"
	fmt.Println(srt2)

	char := 'B'
	fmt.Println(char)

	//Fim strings

	var texto int16
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Null Pointer Exception")
	fmt.Println(erro)

}
