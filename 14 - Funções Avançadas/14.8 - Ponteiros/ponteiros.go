package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	//!  Sem ponteiro: a variavel "numero" continua sendo ela após a execução da função "numero Invertido()"
	numero := -20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//! Com ponteiro: a variavel "numero2" muda permanentemente o seu valor após a execução da função "inverterSinalPonteiro()"
	numero2 := 30
	fmt.Println(numero2)
	inverterSinalPonteiro(&numero2)
	fmt.Println(numero2)

}
