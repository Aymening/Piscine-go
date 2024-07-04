package piscine

func AlphaCount(s string) int {
	counter := 0
	Strg := []rune(s)
	for _, char := range Strg {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			counter++
		}
	}
	return counter
}
