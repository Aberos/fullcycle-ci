package main

import "testing"

func TestSoma(t *testing.T) {
	valorEsperado := 30
	total := Soma(15, 15)

	if total != valorEsperado {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, valorEsperado)
	}
}
