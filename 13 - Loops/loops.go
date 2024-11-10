package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i <= 10; i++ {
		//time.Sleep(time.Second)
		fmt.Println(i)
	}

	nomes := []string{"Rita", "Eduardo", "Renan"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//? Iterar com Strings

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	//? Iterar com map

	usuario := map[string]string{
		"nome":      "Renan",
		"Sobrenome": "Pereira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//? Loop infinito

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
