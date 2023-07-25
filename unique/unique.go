package main

import (
	"fmt"
)

// Função para verificar se um slice contem determinado elemento
func contains(slice []int, elemento int) bool {
	for _, a := range slice {
		if a == elemento {
			return true
		}
	}
	return false
}

func main() {
	//Criando slice enunciado
	slice := [10]int{98, 76, 68, 76, 76, 48, 73, 16, 16, 99}
	fmt.Println(slice)

	//Adicionando elementos(sem repetí-los) no novo slice
	novoSlice := []int{}
	for _, elemento := range slice {

		if !contains(novoSlice, elemento) {
			novoSlice = append(novoSlice, elemento)
		}
	}

	fmt.Println(novoSlice)
}
