package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida das Américas")
	fmt.Println(tipoEndereco)
}
