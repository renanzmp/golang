package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

// ! Adaptado para números (Eu que fiz)
func numerar(_ int, canal2 chan int) {

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		canal2 <- i
		time.Sleep(time.Second)
	}

	close(canal2)
}

func main() {

	var canal = make(chan string)

	go escrever("Olá, Mundo!", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim da escrita")

	fmt.Println("-------------")

	//!
	i := 0
	var canal2 = make(chan int)

	go numerar(i, canal2)

	for mensagem := range canal2 {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}
