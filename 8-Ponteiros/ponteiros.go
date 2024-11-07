package main

import "fmt"

func main() {

	var variavel1 int = 100
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1 += 50

	fmt.Println(variavel1, variavel2)

	//! PONTEIRO É UMA REFERENCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int

	variavel3 = 10
	ponteiro = &variavel3

	variavel3 += 20

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
