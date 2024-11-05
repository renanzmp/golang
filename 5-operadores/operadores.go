package main

import "fmt"

func main() {

	//* ARITIMÉTICOS

	soma := 5 + 3
	sub := 5 - 3
	multi := 4 * 3
	div := 6 / 2
	restoDiv := 6 % 2

	fmt.Println(soma, sub, multi, div, restoDiv)

	//! Só é possível somar numeros do mesmo tipo (int8, int8; int16, int16; etc)
	var num1 int16 = 20
	var num2 int16 = 12

	soma2 := num1 + num2
	fmt.Println(soma2)

	//* ATRIBUIÇÃO

	var variavel1 string = "Olá, mundo!"
	variavel2 := "Salve, tropou"

	fmt.Println(variavel1, variavel2)

	//* OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//* OPERADORES LÓGICOS

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)

	//? Invertendo os valores
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//* OPERADORES UNITÁRIOS

	numero := 10
	numero++
	numero += 15 //numero = numero + 15

	fmt.Println(numero)

	numero--
	numero -= 12

	fmt.Println(numero)

}
