package main

import "fmt"

//!Função declarada antes da main
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//? Função com vários retornos
func calculosMatematicos(n1, n2 int64) (soma int64, sub int64) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	//* Formas de declarar uma função
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Função f!")

	resultado := f("Função f de novo")
	println(resultado)

	//* Função com vários retornos
	resultadoSoma, resultadoSub := calculosMatematicos(30, 20)

	fmt.Println(resultadoSoma, resultadoSub)

	//? Se eu quiser retornar somente um retorno

	somaResultado, _ := calculosMatematicos(30, 20)

	fmt.Println(somaResultado)

}
