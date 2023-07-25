package main

import (
	"fmt"
)

// Criando struct Estado
type Estado struct {
	Nome      string
	Sigla     string
	Capital   string
	Populacao int
	Regiao    string
}

func main() {
	//Instanciando alguns estados
	var estado1 Estado
	estado1 = Estado{"Rio Grande do Norte", "RN", "Natal", 3400000, "Nordeste"}

	var estado2 Estado
	estado2 = Estado{"São Paulo", "SP", "São Paulo", 12000000, "Sudeste"}

	var estado3 Estado
	estado3 = Estado{"Pernambuco", "PE", "Recife", 9300000, "Nordeste"}

	var estado4 Estado
	estado4 = Estado{"Acre", "AC", "Rio Branco", 790000, "Norte"}

	//Criando mapa
	var estados = map[string]Estado{
		"Rio Grande do Norte": estado1,
		"São Paulo":           estado2,
		"Pernambuco":          estado3,
		"Acre":                estado4,
	}

	//Parâmetro que será usado para buscar estados no mapa a partir da key
	var param string = "Rio Grande do Norte"

	fmt.Printf("Qual a capital do %s?\n", param)
	fmt.Println("-", estados[param].Capital)
	fmt.Println("")

	param = "Pernambuco"

	fmt.Printf("Qual a população de %s?\n", param)
	fmt.Println("-", estados[param].Populacao)
	fmt.Println("")

	param = "São Paulo"

	fmt.Printf("Qual a sigla %s?\n", param)
	fmt.Println("-", estados[param].Sigla)
	fmt.Println("")

	param = "Acre"

	fmt.Printf("Em que região o %s está localizado?", param)
	fmt.Println("-", estados[param].Regiao)
	fmt.Println("")
}
