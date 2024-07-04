package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	length := 0
	for i := range a {
		length = i + 1
	}

	for i := 0; i < length-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			descending = false
		} else if f(a[i], a[i+1]) < 0 {
			ascending = false
		}
	}
	return ascending || descending
}
