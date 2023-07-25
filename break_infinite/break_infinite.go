package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a bool = true
	seed := rand.NewSource(time.Now().UnixNano())

	for i := 1; a; i++ {
		random := rand.New(seed)
		numero := random.Intn(10000)
		fmt.Println(numero)

		if numero%42 == 0 {
			fmt.Println("Fim após", i, "iterações.")
			a = false
		}
	}
}
