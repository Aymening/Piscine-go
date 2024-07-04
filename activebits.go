package piscine

func ActiveBits(n int) int {
	count := 0
	s := 0
	for n != 0 {
		s = n % 2
		n = n / 2
		if s == 1 {
			count++
		}
	}
	return count
}
