package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Zumbi",
		"sobrenome": "Palmares",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Ceboso",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "furdúncio",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Cancer",
	}

	fmt.Println(usuario2)
}
