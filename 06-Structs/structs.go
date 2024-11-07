package main

import "fmt"

type endereco struct {
	rua    string
	numero uint16
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo Structs")

	//? Formas de declarar os valores
	//* 1° forma
	var u usuario
	u.nome = "Renan"
	u.idade = 19
	fmt.Println(u)

	endereco := endereco{"Avenida Geremário Dantas", 657}

	//* 2° forma
	usuario2 := usuario{"Renan", 19, endereco}
	fmt.Println(usuario2)

	//* 3° forma (Não tenho todas as informações suficientes)
	usuario3 := usuario{idade: 20}
	fmt.Println(usuario3)

}
