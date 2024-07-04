package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 31 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		result *= RecursiveFactorial(nb-1) * nb
	}
	return result
}
