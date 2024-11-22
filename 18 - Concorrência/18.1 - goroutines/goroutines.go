package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	//? go - O programa não espera a linha de código terminar de ser executada para executar a próxima ( Caso eu não tivesse posto o "go", ele nunca iria para a linha de baixo, escrevendo "Olá mundo" pra sempre)
	go escrever("Olá, mundo")
	escrever("Salveee")

}
