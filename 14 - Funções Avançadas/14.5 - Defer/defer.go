package main

import "fmt"

func alunoAprovado(n1, n2 float64) string {
	defer fmt.Println("Média calculada. Resultado será retornado para você")
	fmt.Println("Entrando na função para calcular a sua média")

	media := (n1 + n2) / 2

	if media >= 7 {
		return ("O aluno foi aprovado")
	} else {
		return ("O aluno foi reprovado")
	}
}

func main() {

	//! Defer adia uma ação até o ultimo momento possível da aplicação
	// defer fmt.Println("Print1")
	// fmt.Println("Print2")

	fmt.Println(alunoAprovado(5, 8))

}
