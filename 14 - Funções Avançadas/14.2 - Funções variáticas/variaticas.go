package main

import "fmt"

func somarNumeros(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	soma := somarNumeros(12, 20, 4, 90)

	fmt.Println(soma)

	escrever("Olá, Mundo", 40, 1, 0, 12)

}
