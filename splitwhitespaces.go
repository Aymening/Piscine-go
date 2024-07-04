package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			str = str + string(s[i])
		}
		if (s[i] == 32 || i == len(s)-1) && (str != "") {
			arr = append(arr, str)
			str = ""
		}
	}
	return arr
}
