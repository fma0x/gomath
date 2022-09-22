package finance

import "math"

// Returns payment amount for financing with fixed installments
func Pmt(amount, rate, months float64) float64 {
	i := rate / 100
	pmt := (amount * ((i * math.Pow((1+i), months)) / (math.Pow((1+i), months) - 1)))
	return math.Round(pmt*100) / 100
}
