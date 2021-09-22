package main

import "testing"

func TestPerimetroRetangulo(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := PerimetroRetangulo(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {

	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
		//esperado float64
	}{
		//{forma: Retangulo{largura: 12, altura: 6}, esperado: 72.0},
		{nome: "Retangulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Circulo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 72.0},
		// {Retangulo{12, 6}, 72.0},
		// {Circulo{10}, 314.1592653589793},
		// {Triangulo{12, 6}, 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}

	// for _, tt := range testesArea {
	// 	resultado := tt.forma.Area()
	// 	if resultado != tt.esperado {
	// 		t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
	// 	}
	// }
}

// func TestArea(t *testing.T) {

// 	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
// 		t.Helper()
// 		resultado := forma.Area()

// 		if resultado != esperado {
// 			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
// 		}
// 	}

// 	t.Run("área retângulo", func(t *testing.T) {
// 		retangulo := Retangulo{12.0, 6.0}
// 		verificaArea(t, retangulo, 72.0)
// 	})

// 	t.Run("área círculo", func(t *testing.T) {
// 		circulo := Circulo{10.0}
// 		verificaArea(t, circulo, 314.1592653589793)
// 	})
// }
