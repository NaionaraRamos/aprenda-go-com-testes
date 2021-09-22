package main

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}

// Codecov
func TestAdd(t *testing.T) {
	if got, want := Adiciona(1, 2), 3; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
