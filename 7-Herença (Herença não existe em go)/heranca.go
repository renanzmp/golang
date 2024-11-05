package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenone string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso string
	turno string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Renan", "Pereira", 19, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciencia da Computação", "noite"}
	fmt.Println(e1)

	fmt.Println(e1.nome)

}
