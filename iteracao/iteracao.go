package main

import "fmt"

//const quantidadeRepeticoes = 5

func Repetir(caractere string, quantidadeRepeticoes int) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

func main() {
	caracteres := Repetir("a", 5)
	fmt.Println(caracteres)
}
