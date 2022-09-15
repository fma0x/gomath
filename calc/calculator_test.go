package calc

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Soma se digito", func(t *testing.T) {
		got := Sum(3, 2)
		if got != 5 {
			t.Errorf("Soma(3, 2) = %v; want 5", got)
		}
	})

	t.Run("Soma com digito", func(t *testing.T) {
		got := Sum(2.5, 2.5)
		if got != 5 {
			t.Errorf("Soma(2.5, 2.5) = %v; want 5", got)
		}
	})
}
