package main

import "testing"

func TestBusca(t *testing.T) {
	//dicionario := map[string]string{"teste": "apenas um teste"}
	//dicionario := make(map[string]string)

	dicionario := Dicionario{"teste": "apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func TestAdiciona(t *testing.T) {

	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "apenas um teste"
		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("Deveria ter encontrado palavra adicionada: ", nil)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, definicao)
	}
}

func TestUpdate(t *testing.T) {

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "apenas um teste"
		novaDefinicao := "nova definicao"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)
		comparaErro(t, err, ErrPalavraInexistente)
	})
}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definicao de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}
