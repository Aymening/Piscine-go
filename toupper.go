package piscine

func ToUpper(s string) string {
	str := []rune(s)

	for i := 0; i < len(s); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] -= 32
		}
	}
	return string(str)
}
