package testes

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0.00
	}
	if amount >= 1000 && amount <= 20000 {
		return 10.00
	}
	if amount >= 20000 {
		return 20.00
	}
	return 5.00

}
