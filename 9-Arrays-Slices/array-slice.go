package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "pos 1"
	fmt.Println(array1)

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(array3)

	// Fim array

	// Slices

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array2[1:3]

	fmt.Println(slice2)

	array2[1] = "Posição alterada"

	fmt.Println(slice2)

	// Array Interno

	fmt.Println("-----------------------------------")

	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)

	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
