package piscine

func TrimAtoi(s string) int {
	boo := true

	str := []rune(s)

	result := 0
	sign := 1
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if str[i] == '-' && boo == true {
			boo = false
			sign = -1

		}
		if str[i] < '0' || str[i] > '9' {
			continue
		} else {
			boo = false
			result = result*10 + int(str[i]-'0')
		}
	}
	return result * sign
}
