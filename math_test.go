package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 1)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: %d / Resultado esperado: %d", total, 30)
	}
}
