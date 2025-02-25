package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 31 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}
