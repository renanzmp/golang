package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	/*//! Esse "for" faz o canal 1 depender dos 2 segundos do canal 2 para escrever no terminal
	for {
		mensagem1 := <-canal1
		fmt.Println(mensagem1)
		mensagem2 := <-canal2
		fmt.Println(mensagem2)
	}*/

	for {
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}

}
