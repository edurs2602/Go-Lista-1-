package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 { //função utilizada para converter celsius em fahrenheit
	fahrenheit := (celsius * 1.8) + 32

	return fahrenheit
}

func drawTable(celsius, fahrenheit float64) string { //função para pegar os valores de celsius e fahrenheit e printar

	fmt.Printf("|   %v |   %v |\n", celsius, fahrenheit)

	return ""
}

func main() {
	initialValue := -40.0 //valor inicial dado pelo enunciado da questão
	lastValue := 100.0    //valor final dado pelo enunciado da questão

	fmt.Println("=================")
	fmt.Println("|    ºC |    ºF |")
	fmt.Println("=================")

	for initialValue <= lastValue { //for loop para verificar se o valor inicial é menor que o final
		temperature := drawTable(initialValue, celsiusToFahrenheit(initialValue)) //passa o valor de initialValue para a função drawTable printar
		fmt.Println(temperature)
		initialValue += 5.0

	}

	fmt.Println("=================")

}
