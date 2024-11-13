package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade > 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Renan", 18}
	usuario1.salvar()

	usuario2 := usuario{"Pedro", 22}
	fmt.Println(usuario2.maiorDeIdade())

	usuario3 := *&usuario{"Marina", 19}
	usuario3.fazerAniversario()
	fmt.Println(usuario3)
	usuario3.fazerAniversario()
	usuario3.fazerAniversario()
	fmt.Println(usuario3)
}
