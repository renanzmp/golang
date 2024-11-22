package main

import "fmt"

func main() {

	//? O posterior à virgula é a capaidade que o canal tem de guardar informações
	canal := make(chan string, 3)

	canal <- "Olá, Mundo!"
	canal <- "Salveee"
	canal <- "Opaa"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
	//! Se eu tentasse criar uma "mensagem 4, daria erro pois a capacidade máxima é 3"

}
