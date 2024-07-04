package piscine

func NRune(s string, n int) rune {
	Slice := []rune(s)
	if n < 1 || n > len(Slice) {
		return 0
	} else {
		return Slice[n-1]
	}
}
