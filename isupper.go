package piscine

func IsUpper(s string) bool {
	str := []rune(s)

	for i := 0; i < len(s); i++ {
		if str[i] < 'A' || str[i] > 'Z' {
			return false
		}
	}
	return true
}
