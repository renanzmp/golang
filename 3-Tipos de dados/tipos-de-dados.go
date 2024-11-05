package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		!Números inteiros
		?Tipos de int:
		*int8
		*int16
		*int32
		*int64
	*/

	var numero1 int = -331200
	fmt.Println(numero1)

	var numero2 uint = 2222
	fmt.Println(numero2)

	//!alias (apelidos)
	//?int32 = rune
	var numero3 rune = 300000
	fmt.Println(numero3)

	//?int8 = byte

	var numero4 byte = 127
	fmt.Println(numero4)

	/*
		!Números reais
		?Tipos:
		*float32
		*float64
	*/

	var numeroReal1 float32 = 328932.42
	fmt.Println(numeroReal1)

	numeroReal2 := 487.9328
	fmt.Println(numeroReal2)

	//!Strings

	var str string = "Olá, mundo!"
	fmt.Println(str)

	//? Char (Te dá o número da Tabela ASCII do caracter)
	char := 'B'
	fmt.Println(char)

	//! Valor 0

	var num int16
	fmt.Println(num)

	var booleano1 bool = true
	var booleano2 bool //Valor 0 do booleano = false
	fmt.Println(booleano1)
	fmt.Println(booleano2)

	var erro1 error
	fmt.Println(erro1) //Valor 0 do erro = nil (null)

	//? Passando um valor para o erro
	var erro2 error = errors.New("Erro detectado")
	fmt.Println(erro2)

}
