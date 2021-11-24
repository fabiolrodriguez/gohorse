package main

import "fmt"

func main() {
	var variavel1 string = "teste"
	variavel2 := "teste2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "dsadsadsa"
		variavel4 string = "dsadsadasdsa"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var 5", "var 6"

	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}
