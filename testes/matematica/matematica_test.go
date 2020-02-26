package matematica

// O arquivo de teste termina em _test

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

// Os métodos de teste começão com Test
func TestMatematica(t *testing.T) {
	// Os métodos de teste precisam receber como parametro t *testing.T
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
