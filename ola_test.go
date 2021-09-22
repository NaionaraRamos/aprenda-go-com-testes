package main

import "testing"

// func TestOla(t *testing.T) {
// 	resultado := Ola()
// 	esperado := "Fala, planeta!"

// 	if resultado != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
// 	}
// }

func TestOla(t *testing.T) {

	// resultado := Ola("Nai")
	// esperado := "Fala, Nai"

	// if resultado != esperado {
	// 	t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	// }

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("cumprimentar as pessoas", func(t *testing.T) {
		resultado := Ola("Nai", "")
		esperado := "Fala, Nai"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Fala, planeta' quando uma string vazia é passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Fala, planeta"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "Espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Elodie", "Francês")
		esperado := "Bonjour, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("testando em espanhol", func(t *testing.T) {
		resultado := Ola("Nai", "Espanhol")
		esperado := "Hola, Nai"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
