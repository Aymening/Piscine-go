package piscine

func IsPrintable(s string) bool {
	for i := range s {
		if s[i] > 126 || s[i] < 32 {
			return false
		}
	}
	return true
}
