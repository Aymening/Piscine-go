package piscine

func Unmatch(a []int) int {
	index := 0
	for _, char := range a {
		index = 0
		for _, char2 := range a {
			if char2 == char {
				index++
			}
		}
		if index%2 != 0 {
			return char
		}
	}
	return -1
}
