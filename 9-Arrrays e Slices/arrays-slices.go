package main

import "fmt"

func main() {

	//! Array tem tamanho fixo. Não consigo adicionar informação posteriormente
	var array1 [5]int
	array1[0] = 10
	array1[3] = 2
	fmt.Println(array1)

	array2 := [4]string{"item1", "item2", "item3", "item4"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//! Slices pode adicionar mais informação posteriormente dependendo da necessidade

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = append(slice1, 25)
	fmt.Println(slice1)
	slice1 = append(slice1, 444)
	fmt.Println(slice1)

	//? Slices pode ser usado para pegar uma "fatia" de uma array

	array4 := [...]string{"Maçã", "Pera", "Abacaxi", "Cenoura", "Cebola", "Morango"}

	slice2 := array4[1:4]
	fmt.Println(slice2)

	array4[2] = "Melancia"

	fmt.Println(slice2)

	//* Arrays internos
	fmt.Println("----------------")

	//tipo, tamanho, capacidade (Desnecessário colocar a capacidade)
	slice4 := make([]int, 10, 11)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("----------")

	slice4 = append(slice4, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("----------")

	//? Sempre que eu estourar a capacidade de um Slice, ele dobra, ou seja, ele tem tamanho infinito
	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
