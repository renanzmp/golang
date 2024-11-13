package main

import "fmt"

func recuperarAplicacao() {
	if r := recover(); r != nil {
		fmt.Println("Aplicação recuperada com sucesso")
	}
}

func AlunoAprovado(n1, n2 float64) string {
	defer recuperarAplicacao()
	media := (n1 + n2) / 2
	if media > 7 {
		return "Você passou"
	} else if media < 7 {
		return "Você reprovou"
	}

	panic("A média é 6!")
}

func main() {

	aluno1 := AlunoAprovado(5, 9)

	fmt.Println(aluno1)
	fmt.Println("Pós execução")

}
