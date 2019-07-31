package main

import "fmt"

func imprimirMayor(v1 int, v2 int) {
	if v1>v2 {
		fmt.Print("El mayor es ", v1)
	} else {
		fmt.Print("El mayor es ", v2)
	}
}

func main() {
	var valor1,valor2 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&valor2)
	imprimirMayor(valor1, valor2)
}
