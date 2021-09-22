package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func testContagem(t *testing.T) {

	t.Run("pausa antes de cada impressão", func(t *testing.T) {

		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})

	t.Run("imprime 3 até Vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		//sleeperSpy := &SleeperSpy{}

		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `
		3
		2
		1
		Vai!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

		// if sleeperSpy.Chamadas != 4 {
		// 	t.Errorf("Não houve chamadas suficientes do sleeper. Esperado 4, resultado %d", sleeperSpy.Chamadas)
		// }
	})
}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second
	tempoSpy := &TempoSpy{}

	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

const pausa = "pausa"
const escrita = "escrita"

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}
