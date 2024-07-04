package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if (str[i] < 'a' || str[i] > 'z') && (str[i] < 'A' || str[i] > 'Z') && (str[i] < '0' || str[i] > '9') {
			return false
		}
	}
	return true
}
