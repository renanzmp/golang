package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//? IF INIT

	if outroNumero := 5; outroNumero > 0 {
		fmt.Println("Número maior que 0")
	} else {
		fmt.Println("Número menor ou igual a 0")
	}

	//! Variavel fica limitada ao escopo do if else
	//fmt.Println(outroNumero)
}
