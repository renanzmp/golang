package main

import "fmt"

//! Aceita tudo

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica(1)
	generica("Olá, mundo!")
	generica(true)
}
