package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("loops")

	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Kiko", "Zinho", "Branco"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonard",
		"Sobrenome": "Nimoy",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
