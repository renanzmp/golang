package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {

	enderecosValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmMinusculo := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmMinusculo, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range enderecosValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}
