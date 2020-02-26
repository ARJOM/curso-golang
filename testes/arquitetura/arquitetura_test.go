package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Significa que o teste pode ser feito de forma concorrente, aumentando o desempenho
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
