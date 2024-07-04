package piscine

func IsLower(s string) bool {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] < 'a' || str[i] > 'z' {
			return false
		}
	}
	return true
}
