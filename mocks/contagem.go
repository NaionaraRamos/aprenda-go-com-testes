package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// func Contagem(saida *bytes.Buffer) {
// 	fmt.Fprint(saida, "3")
// }

type Sleeper interface {
	Pausa()
	//Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	//fmt.Fprint(saida, "3")

	for i := inicioContagem; i > 0; i-- {
		//time.Sleep(1 * time.Second)
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	// for i := inicioContagem; i > 0; i-- {
	// 	fmt.Fprintln(saida, i)
	// }

	//time.Sleep(1 * time.Second)
	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
