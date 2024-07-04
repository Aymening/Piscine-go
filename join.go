package piscine

func Join(strs []string, sep string) string {
	var str string
	for in, char := range strs {
		str += char
		if in != len(strs)-1 {
			str += sep
		}
	}
	return str
}
