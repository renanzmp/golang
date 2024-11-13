package main

import "fmt"

//? INIT é uma função que sera inicializada antes da função main

func init() {
	fmt.Println("Executando da função init")
}

func main() {
	fmt.Println("Executando da função main")
}

func init() {
	fmt.Println("Executando da função init")
}
