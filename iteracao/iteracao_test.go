package main

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 10)
	esperado := "aaaaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", b.N)
	}
}

func ExampleRepetir() {
	repeticao := Repetir("a", 5)
	fmt.Println(repeticao)
	// Output: aaaaa
}
