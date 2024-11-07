package main

import "fmt"

func main() {

	fmt.Println("Maps")

	//* map[tipo da chave] tipo do valor
	usuario := map[string]string{
		"nome":      "Renan",
		"sobrenome": "Pereira",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	fmt.Println("----------")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Renan",
			"segundo":  "Fábio",
			"último":   "Cristiano",
		},
		"curso": {
			"nome":  "Ciencia da Computação",
			"turno": "Noite",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["Ranks"] = map[string]string{
		"Valorant": "Bronze1",
		"CS":       "Global",
	}

	fmt.Println(usuario2)
}
