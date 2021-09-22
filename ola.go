package main

import (
	"fmt"
)

// func Ola() string {
// 	return "Fala, planeta!"
// }

const espanhol = "Espanhol"
const frances = "Francês"
const prefixoOlaPortugues = "Fala, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

//função pública, com letra maiúscula
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "planeta"
	}

	// if idioma == espanhol {
	// 	return prefixoOlaEspanhol + nome
	// }

	// if idioma == frances {
	// 	return prefixoOlaFrances + nome
	// }

	//return prefixoOlaPortugues + nome

	// prefixo := prefixoOlaPortugues

	// switch idioma {
	// case espanhol:
	// 	prefixo = prefixoOlaEspanhol
	// case frances:
	// 	prefixo = prefixoOlaFrances
	// }

	// return prefixo + nome

	return prefixoDeSaudacao(idioma) + nome
}

//função privada, com letra minúscula
func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	//fmt.Println(Ola())
	fmt.Println(Ola("Nai", ""))

	//injecaoDependencias.cumprimenta.
	//Cumprimenta(os.Stdout, "Mel")
}
