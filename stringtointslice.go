package piscine

func StringToIntSlice(str string) []int {
	var s []int
	for _, char := range str {
		s = append(s, int(char))
	}
	return s
}
