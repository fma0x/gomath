package finance

import "testing"

// Teste comparado com o resultado da calculadora do Banco Central
// https://www3.bcb.gov.br/CALCIDADAO/publico/calcularFinanciamentoPrestacoesFixas.do
func TestPmt(t *testing.T) {
	got := Pmt(1000.00, 2, 12)
	expect := 94.56
	if got != expect {
		t.Errorf("Valor de pmt experado de R$ 94.56, recebido R$ %v", got)
	}
}
